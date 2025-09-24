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
	db *chromem.DB // in-memory database
}

func New() *RAG {
	return &RAG{
		db: chromem.NewDB(),
	}
}

func (rag *RAG) Embed(ctx context.Context, name, model string, hs *heapset.HeapSet) *chromem.Collection {
	col, err := rag.db.GetOrCreateCollection(name, map[string]string{
		"model": model,
	}, chromem.NewEmbeddingFuncOllama(model, ""))

	if err != nil {
		sys.Error(err)
		return nil
	}

	var ds []chromem.Document

	hs.Each(func(_ int, h *heap.Heap) {
		if h.Type != types.Agent {
			for _, str := range *h.FMap() {
				ds = append(ds, chromem.Document{
					ID:      fmt.Sprintf("%s:%d", h.Title, str.Nr),
					Content: fmt.Sprintf("line %d: %s\n", str.Nr, str.Str),
				})
			}
		}
	})

	if len(ds) > 0 {
		err = col.AddDocuments(ctx, ds, runtime.NumCPU())
	}

	if err != nil {
		sys.Error(err)
		return nil
	}

	return col
}

func (rag *RAG) Query(ctx context.Context, query string, col *chromem.Collection) string {
	res, err := col.Query(ctx, query, col.Count(), nil, nil)

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
