package openaiclient

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// CreateAzureOpenAIClient initializes a new OpenAI client with the given API key, endpoint, and model name.
func CreateAzureOpenAIClient(apiKey, endpoint, modelName string) *openai.Client {
	config := openai.DefaultAzureConfig(apiKey, endpoint, modelName)
	return openai.NewClientWithConfig(config)
}

// ChatCompletion takes an OpenAI client, a prompt, and a model name as input, and returns a formatted
// response string based on the AI-generated completion or an error if the request fails.
func ChatCompletion(client *openai.Client, prompt string, modelName string) (string, error) {
	// Create a chat completion request with the given model name and user prompt
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: modelName,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	// If there's an error, return an empty string and the error
	if err != nil {
		return "", err
	}

	// Format the response string with the model name and AI-generated completion
	return fmt.Sprintf("Model: %s\n\nResponse: %s", modelName, resp.Choices[0].Message.Content), nil
}
