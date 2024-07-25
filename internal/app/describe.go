// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

package app

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"github.com/DataDog/documentor/internal/ai"
	"github.com/DataDog/documentor/internal/ai/anthropic"
	"github.com/DataDog/documentor/internal/ai/openai"
	"github.com/DataDog/documentor/internal/errno"
	"github.com/DataDog/documentor/internal/prompt"
	"github.com/DataDog/documentor/internal/validate"
	"github.com/urfave/cli/v2"
)

const (
	// ErrEmptyDescribeInput is the error message when the describe command is
	// invoked without any input.
	ErrEmptyDescribeInput xerrors.Error = "missing image file to describe"

	// ErrTooManyImages is the error message when the describe command is
	// invoked with more than one input.
	ErrTooManyImages xerrors.Error = "too many image files to describe; please provide only one"

	// ErrInvalidImageType is the error message when the describe command is
	// invoked with an invalid image file type.
	ErrInvalidImageType xerrors.Error = "invalid image file type; please provide a PNG, JPG, JPEG, or GIF file"

	// ErrInvalidProvider is the error message when the describe command is
	// invoked with an invalid AI provider.
	ErrInvalidProvider xerrors.Error = "invalid AI provider; please refer to the documentation for a list of valid providers"
)

// DescribeAction is the action to perform when the describe command is invoked.
func DescribeAction(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		return errno.New(errno.ExitUsage, ErrEmptyDescribeInput)
	}

	if ctx.Args().Len() > 1 {
		return errno.New(errno.ExitUsage, ErrTooManyImages)
	}

	var (
		key         = ctx.String("key")
		model       = ctx.String("model")
		provider    = ctx.String("provider")
		context     = ctx.String("context")
		temperature = ctx.Float64("temperature")
		filename    = ctx.Bool("filename")
		file        = ctx.Args().Get(0)
		client      ai.Provider
	)

	if !validate.Key(key) {
		return errno.New(errno.ExitUnauthorized, ErrInvalidAPIKey)
	}

	if !validate.Filetype(file, []string{"png", "jpg", "jpeg", "gif"}) {
		return errno.New(errno.ExitInvalidInput, ErrInvalidImageType)
	}

	provider = strings.ToLower(provider)
	provider = strings.TrimSpace(provider)

	switch provider {
	case ai.ProviderOpenAI:
		client = openai.NewClient(key)
	case ai.ProviderAnthropic:
		client = anthropic.NewClient(key)

		if model == openai.DefaultModel {
			model = anthropic.DefaultModel
		}
	default:
		return errno.New(errno.ExitInvalidInput, ErrInvalidProvider)
	}

	data, err := os.ReadFile(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errno.New(errno.ExitNotFound, err)
		}

		if errors.Is(err, os.ErrPermission) {
			return errno.New(errno.ExitPermission, err)
		}

		return errno.New(errno.ExitIO, err)
	}

	userPrompt := "Please generate an SEO-optimized alt text for the attached image."

	if filename {
		userPrompt += " Include a SEO-optimized filename as well."
	}

	req := ai.NewRequestWithImage(data, model, context, userPrompt, prompt.DescribePrompt, float32(temperature))

	err = client.Do(ctx, req)
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return errno.New(errno.ExitTimeout, err)
		}

		return errno.New(errno.ExitAPIError, fmt.Errorf("failed to get response: %w", err))
	}

	return nil
}
