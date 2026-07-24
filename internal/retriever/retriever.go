package retriever

import (
	"context"

	"github.com/Husnain56/rag-from-scratch/internal/embedder"
	"github.com/qdrant/go-client/qdrant"
)

func RetrieveRelevantContent(query string, limit uint64, client *qdrant.Client, collectionName string) ([]string, error) {

	embedding, err := embedder.CreateSingleTextEmbedding(query)

	if err != nil {
		return nil, err
	}

	vec := make([]float32, len(embedding))
	for j, v := range embedding {
		vec[j] = float32(v)
	}

	points, err := QueryPoints(query, limit, client, collectionName, vec)
	if err != nil {
		return nil, err
	}

	content := make([]string, len(points))
	for _, point := range points {
		content = append(content, point.Payload["content"].GetStringValue())
	}
	return content, nil
}

func QueryPoints(query string, limit uint64, client *qdrant.Client, collectionName string, embedding []float32) ([]*qdrant.ScoredPoint, error) {

	points, err := client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: collectionName,
		Query:          qdrant.NewQuery(embedding...),
		Limit:          &limit,
	})

	return points, err
}
