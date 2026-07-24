package embedder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Husnain56/rag-from-scratch/internal/chunker"
)

type EmbedRequest struct {
	Input []string `json:"input"`
	Model string   `json:"model"`
}

type EmbedResponse struct {
	Data []struct {
		Embedding []float64 `json:"embedding"`
	} `json:"data"`
}

func EmbedChunks(chunks []chunker.Chunk) ([]chunker.Chunk, error) {

	texts := make([]string, len(chunks))

	for i := range chunks {
		texts[i] = chunks[i].Content
	}

	embeddings, err := EmbedTexts(texts)

	if err != nil {
		return nil, err
	}

	for i := range chunks {
		chunks[i].Embedding = embeddings[i]
	}
	return chunks, nil
}

func CreateSingleTextEmbedding(query string) ([]float64, error) {

	texts := make([]string, 1)

	texts[0] = query

	embeddings, err := EmbedTexts(texts)

	if err != nil {
		return nil, err
	}

	return embeddings[0], nil

}

func EmbedTexts(texts []string) ([][]float64, error) {

	var request EmbedRequest

	request.Input = append(request.Input, texts...)
	request.Model = "voyage-3-lite"

	jsonData, err := json.Marshal(request)

	if err != nil {
		return nil, fmt.Errorf("failed to convert: %w", err)
	}

	apiUrl := "https://api.voyageai.com/v1/embeddings"
	apiKey := os.Getenv("VOYAGE_API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("Error: VoyageAI_API_KEY environment variable is not set")
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, fmt.Errorf("error building request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var result EmbedResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}
	if len(result.Data) == 0 {
		return nil, fmt.Errorf("no embedding returned")
	}
	embeddings := make([][]float64, len(result.Data))
	for i, d := range result.Data {
		embeddings[i] = d.Embedding
	}

	return embeddings, nil

}
