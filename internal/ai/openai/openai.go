// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

// Package openai provides a client wrapper for the OpenAI API that complies
// with the ai.Provider interface.
package openai

import (
	"errors"
	"fmt"
	"io"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"github.com/DataDog/documentor/internal/ai"
	"github.com/DataDog/documentor/internal/errno"
	"github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

// ErrInvalidRequest is returned when the request provided to the OpenAI API is
// missing both an image and text.
const ErrInvalidRequest xerrors.Error = "invalid request: must provide either an image or text"

// Client represents an OpenAI API client that complies with the ai.Provider
// interface.
type Client struct {
	// ai is the proper OpenAI API client.
	ai *openai.Client
}

// NewClient returns a new Client instance with the given API key.
func NewClient(key string) *Client {
	return &Client{
		ai: openai.NewClient(key),
	}
}

// Compile-time check to ensure Client implements the ai.Provider interface.
var _ ai.Provider = (*Client)(nil)

// Name returns the name of the provider.
func (*Client) Name() string {
	return "OpenAI"
}

// Do performs a single API request to the OpenAI API, returning a response for
// the provided Request and writing said response to ctx.App.Writer as a stream
// of strings.
func (c *Client) Do(ctx *cli.Context, request *ai.Request) error {
	var req openai.ChatCompletionRequest

	switch {
	case request.Image != nil:
		req = NewRequestWithImage(request)
	case request.Text != nil:
		req = NewRequest(request)
	default:
		return ErrInvalidRequest
	}

	resp, err := c.ai.CreateChatCompletionStream(ctx.Context, req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	for {
		text, err := resp.Recv() //nolint:govet // Fixing this is more trouble than it's worth.
		if errors.Is(err, io.EOF) {
			fmt.Fprintf(ctx.App.Writer, "\n")

			break
		}

		if err != nil {
			return errno.New(errno.ExitAPIError, fmt.Errorf("failed to get response: %w", err))
		}

		fmt.Fprintf(ctx.App.Writer, "%s", text.Choices[0].Delta.Content)
	}

	return nil
}
