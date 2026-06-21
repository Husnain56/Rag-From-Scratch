package main

import (
	"fmt"

	"github.com/Husnain56/rag-from-scratch/internal/chunker"
	"github.com/Husnain56/rag-from-scratch/internal/embedder"
	"github.com/Husnain56/rag-from-scratch/internal/loader"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("godotenv error:", err)
	}

	doc, err := loader.LoadDocument("docs/sample.txt")
	if err != nil {
		fmt.Printf("Error loading document: %v\n", err)
		return
	}

	fmt.Printf("Loaded document: %s\n", doc.Metadata["name"])

	chunks := chunker.ChunkDocument(doc)

	fmt.Printf("Total chunks: %d\n", len(chunks))

	// for _, chunk := range chunks {
	// 	fmt.Printf("\n--- Chunk %s ---\n", chunk.Metadata["chunk_index"])
	// 	fmt.Printf("%s\n", chunk.Content)
	// }

	vector, err := embedder.EmbedText(chunks[0].Content)
	if err != nil {
		fmt.Println("Embedding error:", err)
		return
	}

	fmt.Printf("Chunk text: %s\n", chunks[0].Content)
	fmt.Printf("Vector length: %d\n", len(vector))
	fmt.Printf("First 5 values: %v\n", vector[:5])

}
