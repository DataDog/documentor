// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

package app

import (
	"errors"
	"fmt"
	"io"
	"os"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
	"github.com/DataDog/documentor/internal/ai/openai"
	"github.com/DataDog/documentor/internal/errno"
	"github.com/DataDog/documentor/internal/prompt"
	"github.com/DataDog/documentor/internal/validate"
	"github.com/urfave/cli/v2"
)

const (
	// ErrEmptyDraftInput is the error message when the draft command is
	// invoked without any input.
	ErrEmptyDraftInput xerrors.Error = "missing input for draft; please provide a file with your notes"

	// ErrTooManyNotes is the error message when the draft command is
	// invoked with more than one input.
	ErrTooManyNotes xerrors.Error = "too many notes files; please provide only one"

	// ErrInvalidFiletype is the error message when the draft command is
	// invoked with a file that is not a text file.
	ErrInvalidFiletype xerrors.Error = "invalid filetype; please provide a TXT or MD file"
)

// DraftAction is the action to perform when the draft command is invoked.
func DraftAction(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		return errno.New(errno.ExitUsage, ErrEmptyDraftInput)
	}

	if ctx.Args().Len() > 1 {
		return errno.New(errno.ExitUsage, ErrTooManyNotes)
	}

	var (
		key         = ctx.String("key")
		model       = ctx.String("model")
		temperature = ctx.Float64("temperature")
		file        = ctx.Args().Get(0)
	)

	if !validate.Key(key) {
		return errno.New(errno.ExitUnauthorized, ErrInvalidAPIKey)
	}

	if !validate.Filetype(file, []string{"txt", "md"}) {
		return errno.New(errno.ExitInvalidInput, ErrInvalidFiletype)
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
		notes   = xunsafe.BytesToString(data)
		content = "Please write a new technical document based on the notes " +
			"provided below. It's very important that I get a good answer " +
			"as I'm under a LOT of stress at work. I'll tip $500 if you can " +
			"help me out. Here are the notes:\n\n" + notes
		client = openai.NewClient(key)
		req    = openai.NewRequest(content, model, prompt.DraftPrompt, float32(temperature))
	)

	resp, err := client.Do(ctx.Context, req)
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return errno.New(errno.ExitTimeout, err)
		}

		return errno.New(errno.ExitAPIError, fmt.Errorf("failed to get response: %w", err))
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
