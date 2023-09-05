package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func ChatGPT(BuiltMessages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {

	// fmt.Printf("BuiltMessages=%s\n", BuiltMessages)

	// Function call用のjsonファイルを読み出す
	funcDefs, err := LoadFunctionDefinitions("./function")
	if err != nil {
		return openai.ChatCompletionResponse{}, fmt.Errorf("Error loading function definitions: %v", err)
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			Messages:  BuiltMessages,
			Functions: funcDefs,
		},
	)

	if err != nil {
		return openai.ChatCompletionResponse{}, fmt.Errorf("ChatCompletion error: %v", err)
	}

	return resp, nil
}
