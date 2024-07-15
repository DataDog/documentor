// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

package openai

import (
	_ "embed"

	"github.com/sashabaranov/go-openai"
)

//go:embed data/review-prompt.txt
var MarkdownPrompt string

//go:embed data/describe-prompt.txt
var DescribePrompt string

//go:embed data/draft-prompt.txt
var DraftPrompt string

// NewRequest creates a chat completion request with streaming support for the
// OpenAI API given the content of the chat.
func NewRequest(content, systemPrompt string, temperature float32) openai.ChatCompletionRequest {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
		},
		Temperature: temperature,
		Stream:      true,
	}

	return req
}

// NewRequestWithImage creates a chat completion request with streaming support
// for the OpenAI API given a base64 encoded image.
func NewRequestWithImage( //nolint:revive // I really don't feel like creating another function for this.
	image, context, systemPrompt string,
	filename bool,
	temperature float32,
) openai.ChatCompletionRequest {
	if context != "" {
		systemPrompt = systemPrompt + "\n\nContext for the image:\n" + context
	}

	userPrompt := "Please generate an SEO-optimized alt text for the attached image."

	if filename {
		userPrompt += " Include a SEO-optimized filename as well."
	}

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: userPrompt,
			},
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL:    image,
							Detail: openai.ImageURLDetailAuto,
						},
					},
				},
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
		},
		Temperature: temperature,
		Stream:      true,
	}

	return req
}
