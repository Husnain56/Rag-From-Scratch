package main

import (
	"fmt"

	"github.com/Husnain56/rag-from-scratch/internal/generator"
	"github.com/Husnain56/rag-from-scratch/internal/retriever"
	"github.com/Husnain56/rag-from-scratch/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("godotenv error:", err)
	}

	// doc, err := loader.LoadDocument("docs/sample.txt")
	// if err != nil {
	// 	fmt.Printf("Error loading document: %v\n", err)
	// 	return
	// }

	// fmt.Printf("Loaded document: %s\n", doc.Metadata["name"])

	// chunks := chunker.ChunkDocument(doc)

	// fmt.Printf("Total chunks: %d\n", len(chunks))

	// chunks, err = embedder.EmbedChunks(chunks)
	// if err != nil {
	// 	fmt.Println("Error setting embeddings:", err)
	// 	return
	// }

	qdrantClient, err := store.NewClient()
	if err != nil {
		fmt.Println("Error creating Qdrant client:", err)
		return
	}

	content, err := retriever.RetrieveRelevantContent("What is goLang?", 5, qdrantClient, "my_collection")
	if err != nil {
		fmt.Println("Error retrieving relevant content:", err)
		return
	}

	fmt.Printf("%s", content[0])

	query := "What is GoLang?"

	prompt := generator.BuildPrompt(query, content)

	response, err := generator.GenerateResponse(prompt)

	fmt.Printf("%s", response)

	// points, err := store.BuildPoints(chunks)
	// if err != nil {
	// 	fmt.Println("Error building points:", err)
	// 	return
	// }
	// fmt.Printf("Total points: %d\n", len(points))
	// fmt.Printf("First point ID: %v\n", points[0].Id)
	// fmt.Printf("Vector size: %d\n", len(points[0].Vectors.GetVector().Data))

	// err = store.CreateCollection(qdrantClient, "my_collection", uint64(len(chunks[0].Embedding)))

	// err = store.AddPointsToCollection(qdrantClient, "my_collection", points)
	// if err != nil {
	// 	fmt.Println("Upsert error:", err)
	// 	return
	// }
	// fmt.Println("Points upserted successfully!")

}
