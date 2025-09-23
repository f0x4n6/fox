package rag

import (
	"context"
	"runtime"
	"strings"

	"github.com/philippgille/chromem-go"

	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/types"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

const prefix = "search_document: "

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
			if _, ok := rag.idx[h.String()]; !ok {
				docs = append(docs, chromem.Document{
					ID:      h.String(),
					Content: prefix + h.Content(),
					Metadata: map[string]string{
						"source": h.Base,
					},
				})

				rag.idx[h.String()] = types.Nop{}
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
		sb.WriteString(strings.TrimPrefix(r.Content, prefix))
		sb.WriteRune('\n')
	}

	return sb.String()
}
