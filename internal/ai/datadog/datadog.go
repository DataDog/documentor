// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

// Package datadog provides a client wrapper for the Datadog AI proxy API that
// complies with the ai.Provider interface.
package datadog

import (
	"fmt"
	"net/http"
	"time"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"github.com/DataDog/documentor/internal/ai"
	"github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

// ErrInvalidRequest is returned when the request provided to the Datadog proxy
// API is missing both an image and text.
const ErrInvalidRequest xerrors.Error = "invalid request: must provide either an image or text"

// DefaultModel is the default model to use when making requests to the API.
const DefaultModel = "gpt-4o"

// Client represents a Datadog AI proxy API client that complies with the
// ai.Provider interface.
type Client struct {
	// ai is the proper Datadog proxy client.
	ai *openai.Client
}

// NewClient returns a new Client instance with the given email.
func NewClient(endpoint, email string) *Client {
	cfg := openai.DefaultConfig("not-a-real-key")
	cfg.BaseURL = endpoint
	cfg.HTTPClient = &http.Client{
		Transport: &Transport{
			Email: email,
		},
		Timeout: 15 * time.Second,
	}

	return &Client{
		ai: openai.NewClientWithConfig(cfg),
	}
}

// Compile-time check to ensure Client implements the ai.Provider interface.
var _ ai.Provider = (*Client)(nil)

// Name returns the name of the provider.
func (*Client) Name() string {
	return "Datadog"
}

// Do performs a single API request to the Datadog API, returning a response for
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

	resp, err := c.ai.CreateChatCompletion(ctx.Context, req)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	fmt.Fprintf(ctx.App.Writer, "%s\n", resp.Choices[0].Message.Content)

	return nil
}
