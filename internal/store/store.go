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
