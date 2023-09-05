package main

import (
	openai "github.com/sashabaranov/go-openai"
)

func BuildMessages(Message string) (BuiltMessages []openai.ChatCompletionMessage) {

	// fmt.Printf("MessageTest:\n%s\n",Message)

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: `You are a helpful assistant.`,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: Message,
		},
	}
	return messages
}
