package rag

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/philippgille/chromem-go"

	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
	"github.com/cuhsat/fox/internal/pkg/types/smap"
)

type RAG struct {
	db  *chromem.DB // in-memory database
	idx sync.Map    // in-memory index
}

func New() *RAG {
	return &RAG{
		db: chromem.NewDB(),
	}
}

func (rag *RAG) Embed(ctx context.Context, model string, heap *heap.Heap) *chromem.Collection {
	col, err := rag.db.GetOrCreateCollection(heap.String(), map[string]string{
		"model": model,
	}, chromem.NewEmbeddingFuncOllama(model, ""))

	if err != nil {
		sys.Error(err)
		return nil
	}

	heap.FMap().Extern(func(str smap.String) {
		if _, ok := rag.idx.Load(str.Nr); !ok {
			if err := col.AddDocument(ctx, chromem.Document{
				ID:      strconv.Itoa(str.Nr),
				Content: fmt.Sprintf("line %d: %s\n", str.Nr, str.Str),
			}); err == nil {
				rag.idx.Store(str.Nr, struct{}{})
			} else {
				sys.Error(err)
			}
		}
	})

	return col
}

func (rag *RAG) Query(ctx context.Context, query string, col *chromem.Collection) string {
	res, err := col.Query(ctx, query, 10, nil, nil)

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
