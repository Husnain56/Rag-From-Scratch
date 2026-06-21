package chunker

import (
	"fmt"
	"strings"
	"github.com/Husnain56/rag-from-scratch/internal/loader"
)

type Chunk struct {
	Content   string
	Metadata  map[string]string
	Embedding []float64
}

func ChunkDocument(doc loader.Document) []Chunk {
	var chunks []Chunk

	normalized := strings.ReplaceAll(doc.Content, "\r\n", "\n")
	paragraphs := strings.Split(normalized, "\n\n")

	for i, paragraph := range paragraphs {
		trimmed := strings.TrimSpace(paragraph)
		if trimmed == "" {
			continue
		}

		chunk := Chunk{
			Content: trimmed,
			Metadata: map[string]string{
				"source":      doc.Metadata["source"],
				"filename":    doc.Metadata["name"],
				"chunk_index": fmt.Sprintf("%d", i),
			},
		}
		chunks = append(chunks, chunk)
	}

	return chunks
}
