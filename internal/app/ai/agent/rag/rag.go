package rag

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/philippgille/chromem-go"

	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

type RAG struct {
	db  *chromem.DB // in-memory database
	idx types.Set   // in-memory index
}

func New() *RAG {
	return &RAG{
		db:  chromem.NewDB(),
		idx: types.Set{},
	}
}

func (rag *RAG) Embed(name, model string, hs *heapset.HeapSet) *chromem.Collection {
	col, err := rag.db.GetOrCreateCollection(name, map[string]string{
		"model": model,
	}, chromem.NewEmbeddingFuncOllama(model, ""))

	if err != nil {
		sys.Error(err)
		return nil
	}

	var docs []chromem.Document

	hs.Each(func(_ int, h *heap.Heap) {
		if h.Type != types.Agent {
			for _, str := range *h.FMap() {
				id := key(h.Base, str.Nr)

				if _, ok := rag.idx[id]; !ok {
					docs = append(docs, chromem.Document{
						ID:       id,
						Metadata: map[string]string{"file": h.Base},
						Content:  fmt.Sprintf("line %d: %s", str.Nr, str.Str),
					})

					rag.idx[id] = types.Nop{}
				}
			}
		}
	})

	if len(docs) > 0 {
		err = col.AddDocuments(context.Background(), docs, runtime.NumCPU())
	}

	if err != nil {
		sys.Error(err)
		return nil
	}

	return col
}

func (rag *RAG) Query(query string, col *chromem.Collection) string {
	res, err := col.Query(context.Background(), query, col.Count(), nil, nil)

	if err != nil {
		sys.Error(err)
		return ""
	}

	var sb strings.Builder

	for _, r := range res {
		sb.WriteString(r.Content)
		sb.WriteRune('\n')
	}

	return sb.String()
}

func key(f string, l int) string {
	return fmt.Sprintf("%s:%d", f, l)
}
