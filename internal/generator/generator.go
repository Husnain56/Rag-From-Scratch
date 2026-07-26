package generator

import (
	"fmt"
	"strings"
)

func BuildPrompt(query string, content []string) string {

	prompt := fmt.Sprintf(`You are a helpful support assistant. Answer the user's question only
	from the facts provided in the context below. If the answer cannot be found
	in the context below, just respond with "I cannot find this information."
	Do not use outside knowledge or make assumptions.

	Context: %s

	Question: %s`, strings.Join(content, "\n\n"), query)

	return prompt
}
