// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

package openai

import (
	_ "embed"

	"github.com/sashabaranov/go-openai"
)

//go:embed data/prompt.txt
var _systemPrompt string

// NewRequest creates a chat completion request with streaming support for the
// OpenAI API given the content of the chat.
func NewRequest(content string) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				Content: "Please review the following content. It's very " +
					"important that I get a good answer as I'm under a LOT of " +
					"stress at work. I'll tip $500 if you can help me.\n\n" +
					content,
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: _systemPrompt,
			},
		},
		Temperature: 0.8,
		Stream:      true,
	}
}
