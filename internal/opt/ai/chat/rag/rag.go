package rag

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/philippgille/chromem-go"

	"github.com/cuhsat/fox/v3/internal/pkg/types/heap"
	"github.com/cuhsat/fox/v3/internal/pkg/types/smap"
)

type RAG struct {
	db    *chromem.DB // in-memory database
	index sync.Map    // line index
}

func New() *RAG {
	return &RAG{
		db: chromem.NewDB(),
	}
}

func (rag *RAG) Query(ctx context.Context, query string, col *chromem.Collection) string {
	res, err := col.Query(ctx, query, col.Count(), nil, nil)

	if err != nil {
		log.Println(err)
		return ""
	}

	var sb strings.Builder

	for _, r := range res {
		sb.WriteString(r.Content)
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (rag *RAG) Embed(ctx context.Context, model string, heap *heap.Heap) *chromem.Collection {
	col, err := rag.db.GetOrCreateCollection(heap.String(), map[string]string{
		"model": model,
	}, chromem.NewEmbeddingFuncOllama(model, ""))

	if err != nil {
		log.Println(err)
		return nil
	}

	heap.FMap().Extern(func(str smap.String) {
		if _, ok := rag.index.Load(str.Nr); !ok {
			if err := col.AddDocument(ctx, chromem.Document{
				ID:      strconv.Itoa(str.Nr),
				Content: fmt.Sprintf("line %d: %s\n", str.Nr, str.Str),
			}); err == nil {
				rag.index.Store(str.Nr, struct{}{})
			} else {
				log.Println(err)
			}
		}
	})

	return col
}
