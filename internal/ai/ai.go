// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

// Package ai provides an interface for interacting with AI providers.
package ai

import "github.com/urfave/cli/v2"

// Provider represents an LLM provider such as OpenAI, Anthropic, Mistral, etc.
type Provider interface {
	// Name returns the name of the provider.
	Name() string

	// Do performs a single API request to the provider's API, returning a
	// response for the provided Request.
	//
	// Do should attempt to interpret the response from the provider and return
	// it as streaming strings to ctx.App.Writer if successful.
	Do(ctx *cli.Context, req *Request) error
}

// Request represents an HTTP request to an AI provider's API.
type Request struct {
	// Content is the user prompt to send to the AI provider. It should be
	// either a string with the prompt text or a base64-encoded image, depending
	// on the type of request.
	Content string

	// Context is a string that provides additional context for the request. It
	// is only used if Content is a base64-encoded image.
	Context string

	// Model is the full name of the LLM model to use for the request, e.g.
	// "gpt-4o-mini" or "claude-3-5-sonnet-20240620".
	Model string

	// SystemPrompt is a set of instructions for the AI model to follow when
	// generating a response.
	SystemPrompt string

	// UserPrompt is a second set of instructions for the AI model to follow
	// when generating a response.
	//
	// This prompt is sent to the AI provider after the system prompt and can be
	// used to provide additional context or constraints for the response.
	UserPrompt string

	// Temperature is a float between 0 and 1 that controls the randomness of
	// the response. A value of 0 will always return the most likely token,
	// while a value of 1 will sample from the distribution of tokens.
	Temperature float32
}

// NewRequest returns a new Request instance for a text-based user prompt.
func NewRequest(model, content, userPrompt, systemPrompt string, temperature float32) *Request {
	return &Request{
		Content:      content,
		Model:        model,
		SystemPrompt: systemPrompt,
		UserPrompt:   userPrompt,
		Temperature:  temperature,
	}
}

// NewRequestWithImage returns a new Request instance for an user prompt with a
// base64-encoded image attachment.
func NewRequestWithImage(model, image, context, userPrompt, systemPrompt string, temperature float32) *Request {
	if context != "" {
		userPrompt += "/n/n Here is some context for the image: " + context
	}

	return &Request{
		Content:      image,
		Context:      context,
		Model:        model,
		SystemPrompt: systemPrompt,
		UserPrompt:   userPrompt,
		Temperature:  temperature,
	}
}
