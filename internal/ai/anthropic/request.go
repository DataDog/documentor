package anthropic

import (
	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
	"github.com/DataDog/documentor/internal/ai"
	"github.com/liushuangls/go-anthropic/v2"
)

// NewRequest creates a chat completion request with streaming support for the
// Anthropic API given the provided ai.Request object.
func NewRequest(req *ai.Request) anthropic.MessagesStreamRequest {
	text := xunsafe.BytesToString(req.Text)

	return anthropic.MessagesStreamRequest{
		MessagesRequest: anthropic.MessagesRequest{
			Model:       req.Model,
			Temperature: &req.Temperature,
			MaxTokens:   4096,
			System:      req.SystemPrompt,
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
	}
}

// NewRequestWithImage creates a chat completion request with streaming support
// for the Anthropic API given the provided ai.Request object. The req.Content
// field should be a base64-encoded image.
func NewRequestWithImage(req *ai.Request) anthropic.MessagesStreamRequest {
	return anthropic.MessagesStreamRequest{
		MessagesRequest: anthropic.MessagesRequest{
			Model:       req.Model,
			Temperature: &req.Temperature,
			MaxTokens:   4096,
			System:      req.SystemPrompt,
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
	}
}
