package retriever

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func QueryChunks(query string, limit uint64, client *qdrant.Client, collectionName string, embedding []float32) ([]*qdrant.ScoredPoint, error) {

	points, err := client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: collectionName,
		Query:          qdrant.NewQuery(embedding...),
		Limit:          &limit,
	})

	return points, err
}
