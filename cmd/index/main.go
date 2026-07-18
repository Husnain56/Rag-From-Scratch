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

	chunks, err = embedder.EmbedChunks(chunks)
	if err != nil {
		fmt.Println("Error setting embeddings:", err)
		return
	}

}
