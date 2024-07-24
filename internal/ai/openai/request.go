// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

package openai

import (
	"github.com/DataDog/documentor/internal/ai"
	"github.com/sashabaranov/go-openai"
)

// NewTextRequest creates a chat completion request with streaming support for
// the OpenAI API given the provided ai.Request object.
func NewRequest(req *ai.Request) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model:       req.Model,
		Temperature: req.Temperature,
		Stream:      true,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: req.Content,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: req.UserPrompt,
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: req.SystemPrompt,
			},
		},
	}
}

// NewRequestWithImage creates a chat completion request with streaming support
// for the OpenAI API given the provided ai.Request object. The req.Content
// field should be a base64-encoded image.
func NewRequestWithImage(req *ai.Request) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model:       req.Model,
		Temperature: req.Temperature,
		Stream:      true,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: req.UserPrompt,
			},
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL:    req.Content,
							Detail: openai.ImageURLDetailAuto,
						},
					},
				},
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: req.SystemPrompt,
			},
		},
	}
}
