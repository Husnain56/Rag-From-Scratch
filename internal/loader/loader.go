package loader

import (
	"fmt"
	"os"
)

type Document struct {
	Content  string
	Metadata map[string]string
}

func LoadDocument(filePath string) (Document, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return Document{}, fmt.Errorf("failed to read file: %w", err)
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return Document{}, fmt.Errorf("failed to get file info: %w", err)
	}

	return Document{
		Content: string(content),
		Metadata: map[string]string{
			"source": filePath,
			"name":   fileInfo.Name(),
			"size":   fmt.Sprintf("%d", fileInfo.Size()),
		},
	}, nil
}
