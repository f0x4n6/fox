package rag

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/philippgille/chromem-go"

	"github.com/cuhsat/fox/internal/pkg/sys"
	"github.com/cuhsat/fox/internal/pkg/types/heap"
)

type RAG struct {
	db *chromem.DB // in-memory database
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

	for _, str := range *heap.FMap() {
		err := col.AddDocument(ctx, chromem.Document{
			ID:      strconv.Itoa(str.Nr),
			Content: fmt.Sprintf("line %d: %s\n", str.Nr, str.Str),
		})

		if err != nil {
			sys.Error(err)
			return nil
		}
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
