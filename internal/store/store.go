package store

import (
	"context"
	"os"

	"github.com/qdrant/go-client/qdrant"
)

func NewClient() (*qdrant.Client, error) {

	endpoint := os.Getenv("Qdrant_ENDPOINT")
	apiKey := os.Getenv("Qdrant_API_KEY")

	return qdrant.NewClient(&qdrant.Config{
		Host:   endpoint, // e.g. "xxxx.cloud.qdrant.io"
		Port:   6334,     // gRPC port, not 6333 (that's REST)
		APIKey: apiKey,
		UseTLS: true,
	})
}

func CreateCollection(client *qdrant.Client, collectionName string, vectorSize uint64) error {

	exists, err := client.CollectionExists(context.Background(), "my_collection")

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
