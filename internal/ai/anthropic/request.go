// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

package anthropic

import (
	"fmt"

	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
	"github.com/DataDog/documentor/internal/ai"
	"github.com/liushuangls/go-anthropic/v2"
	"github.com/urfave/cli/v2"
)

// NewRequest creates a chat completion request with streaming support for the
// Anthropic API given the provided ai.Request object.
func NewRequest(ctx *cli.Context, req *ai.Request) anthropic.MessagesStreamRequest {
	text := xunsafe.BytesToString(req.Text)

	return anthropic.MessagesStreamRequest{
		MessagesRequest: anthropic.MessagesRequest{
			Model:       req.Model,
			Temperature: &req.Temperature,
			System:      req.SystemPrompt,
			MaxTokens:   4096,
			Stream:      true,
			Messages: []anthropic.Message{
				{
					Role: anthropic.RoleUser,
					Content: []anthropic.MessageContent{
						anthropic.NewTextMessageContent(text),
						anthropic.NewTextMessageContent(req.UserPrompt),
					},
				},
			},
		},
		OnContentBlockDelta: func(data anthropic.MessagesEventContentBlockDeltaData) {
			fmt.Fprintf(ctx.App.Writer, "%s", *data.Delta.Text)
		},
	}
}

// NewRequestWithImage creates a chat completion request with streaming support
// for the Anthropic API given the provided ai.Request object. The req.Content
// field should be a base64-encoded image.
func NewRequestWithImage(ctx *cli.Context, req *ai.Request) anthropic.MessagesStreamRequest {
	return anthropic.MessagesStreamRequest{
		MessagesRequest: anthropic.MessagesRequest{
			Model:       req.Model,
			Temperature: &req.Temperature,
			System:      req.SystemPrompt,
			MaxTokens:   4096,
			Stream:      true,
			Messages: []anthropic.Message{
				{
					Role: anthropic.RoleUser,
					Content: []anthropic.MessageContent{
						anthropic.NewImageMessageContent(anthropic.MessageContentImageSource{
							Type:      "base64",
							MediaType: "image/jpeg",
							Data:      req.Image,
						}),
						anthropic.NewTextMessageContent(req.UserPrompt),
					},
				},
			},
		},
		OnContentBlockDelta: func(data anthropic.MessagesEventContentBlockDeltaData) {
			fmt.Fprintf(ctx.App.Writer, "%s", *data.Delta.Text)
		},
	}
}
