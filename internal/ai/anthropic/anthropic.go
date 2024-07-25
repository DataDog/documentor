// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

// Package anthropic provides a client wrapper for the Anthropic API that
// complies with the ai.Provider interface.
package anthropic

import (
	"fmt"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"github.com/DataDog/documentor/internal/ai"
	"github.com/liushuangls/go-anthropic/v2"
	"github.com/urfave/cli/v2"
)

// ErrInvalidRequest is returned when the request provided to the OpenAI API is
// missing both an image and text.
const ErrInvalidRequest xerrors.Error = "invalid request: must provide either an image or text"

// Client represents an Anthropic API client that complies with the ai.Provider
// interface.
type Client struct {
	// ai is the proper Anthropic API client.
	ai *anthropic.Client
}

// NewClient returns a new Client instance with the given API key.
func NewClient(key string) *Client {
	return &Client{
		ai: anthropic.NewClient(key),
	}
}

// Compile-time check to ensure Client implements the ai.Provider interface.
var _ ai.Provider = (*Client)(nil)

// Name returns the name of the provider.
func (*Client) Name() string {
	return "Anthropic"
}

// Do performs a single API request to the Anthropic API, returning a response for
// the provided Request and writing said response to ctx.App.Writer as a stream
// of strings.
func (c *Client) Do(ctx *cli.Context, request *ai.Request) error {
	var req anthropic.MessagesStreamRequest

	if request.Image != nil {
		req = NewRequestWithImage(ctx, request)
	} else if request.Text != nil {
		req = NewRequest(ctx, request)
	} else {
		return ErrInvalidRequest
	}

	_, err := c.ai.CreateMessagesStream(ctx.Context, req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	fmt.Fprintf(ctx.App.Writer, "\n")

	return nil
}
