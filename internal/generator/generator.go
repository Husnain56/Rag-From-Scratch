package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type DeepSeekRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type DeepSeekResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func BuildPrompt(query string, content []string) string {

	prompt := fmt.Sprintf(`You are a helpful support assistant. Answer the user's question only
	from the facts provided in the context below. If the answer cannot be found
	in the context below, just respond with "I cannot find this information."
	Do not use outside knowledge or make assumptions.

	Context: %s

	Question: %s`, strings.Join(content, "\n\n"), query)

	return prompt
}

func GenerateResponse(prompt string) (string, error) {

	apiUrl := os.Getenv("DeepSeek_API_URL")
	apiKey := os.Getenv("DeepSeek_API_KEY")

	if apiKey == "" {
		return "", fmt.Errorf("Error: VoyageAI_API_KEY environment variable is not set")
	}

	request := DeepSeekRequest{
		Model: "deepseek-chat",
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonData, err := json.Marshal(request)

	if err != nil {
		return "", fmt.Errorf("failed to convert: %w", err)
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))

	if err != nil {
		return "", fmt.Errorf("error building request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	var result DeepSeekResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no response generated")
	}

	return result.Choices[0].Message.Content, nil

}
