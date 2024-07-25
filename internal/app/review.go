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
	// ErrEmptyInput is returned when no file is provided by the user.
	ErrEmptyInput xerrors.Error = "missing file to review"

	// ErrInvalidAPIKey is returned when the user does not provide an API key
	// for the given provider or provides an invalid one.
	ErrInvalidAPIKey xerrors.Error = "missing or invalid API key"

	// ErrTooMuchInput is returned when the user provides more than one file.
	ErrTooMuchInput xerrors.Error = "too many files to review; please provide only one file"
)

// ReviewAction is the main action for the application.
func ReviewAction(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		return errno.New(errno.ExitUsage, ErrEmptyInput)
	}

	if ctx.Args().Len() > 1 {
		return errno.New(errno.ExitUsage, ErrTooMuchInput)
	}

	var (
		key         = ctx.String("key")
		model       = ctx.String("model")
		provider    = ctx.String("provider")
		temperature = ctx.Float64("temperature")
		file        = ctx.Args().Get(0)
		client      ai.Provider
	)

	if !validate.Key(key) {
		return errno.New(errno.ExitUnauthorized, ErrInvalidAPIKey)
	}

	if !validate.Filetype(file, []string{"txt", "md"}) {
		return errno.New(errno.ExitInvalidInput, ErrInvalidFiletype)
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

	var (
		userPrompt = "Please review the following content. It's very " +
			"important that I get a good answer as I'm under a LOT of " +
			"stress at work. I'll tip $500 if you can help me."
		req = ai.NewRequest(data, model, userPrompt, prompt.MarkdownPrompt, float32(temperature))
	)

	err = client.Do(ctx, req)
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return errno.New(errno.ExitTimeout, err)
		}

		return errno.New(errno.ExitAPIError, fmt.Errorf("failed to get response: %w", err))
	}

	return nil
}
