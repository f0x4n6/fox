package rag

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/philippgille/chromem-go"

	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/types/heapset"
)

type RAG struct {
	db *chromem.DB // in-memory database
}

func New() *RAG {
	return &RAG{
		db: chromem.NewDB(),
	}
}

func (rag *RAG) Embed(name, model string, hs *heapset.HeapSet) *chromem.Collection {
	fn := chromem.NewEmbeddingFuncOllama(model, "")

	col, err := rag.db.GetOrCreateCollection(name, nil, fn)

	// TODO: Clear collection before

	if err != nil {
		sys.Error(err)
		return nil
	}

	var docs []chromem.Document

	hs.Each(func(i int, h *heap.Heap) {
		for j, str := range *h.FMap() {
			docs = append(docs, chromem.Document{
				ID:       fmt.Sprintf("%d:%d", i, j),
				Metadata: map[string]string{"path": h.Base},
				Content:  fmt.Sprintf("line %d: %s", str.Nr, str.Str),
			})
		}
	})

	err = col.AddDocuments(context.Background(), docs, runtime.NumCPU())

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
