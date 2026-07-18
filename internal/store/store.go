package store

import (
	"context"
	"os"

	"github.com/Husnain56/rag-from-scratch/internal/chunker"
	"github.com/qdrant/go-client/qdrant"
)

func NewClient() (*qdrant.Client, error) {

	endpoint := os.Getenv("Qdrant_ENDPOINT")
	apiKey := os.Getenv("Qdrant_API_KEY")

	return qdrant.NewClient(&qdrant.Config{
		Host:   endpoint,
		Port:   6334,
		APIKey: apiKey,
		UseTLS: true,
	})
}

func CreateCollection(client *qdrant.Client, collectionName string, vectorSize uint64) error {

	exists, err := client.CollectionExists(context.Background(), collectionName)

	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	return client.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName: collectionName,
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     vectorSize,
			Distance: qdrant.Distance_Cosine,
		}),
	})

}

func buildPoints(chunks []chunker.Chunk) ([]*qdrant.PointStruct, error) {

	points := make([]*qdrant.PointStruct, len(chunks))

	for i, chunk := range chunks {

		vec := make([]float32, len(chunk.Embedding))
		for j, v := range chunk.Embedding {
			vec[j] = float32(v)
		}

		points[i] = &qdrant.PointStruct{
			Id:      qdrant.NewIDNum(uint64(i)),
			Vectors: qdrant.NewVectors(vec...),
			Payload: qdrant.NewValueMap(map[string]any{
				"content":  chunk.Content,
				"metadata": chunk.Metadata,
			}),
		}
	}
	return points, nil
}

func addPointsToCollection(client *qdrant.Client, collectionName string, points []*qdrant.PointStruct) error {

	_, err := client.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points:         points,
	})

	if err != nil {
		return err
	}

	return nil

}
