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
	"github.com/DataDog/documentor/internal/openai"
	"github.com/urfave/cli/v2"
)

const (
	// ErrEmptyInput is returned when no file is provided by the user.
	ErrEmptyInput xerrors.Error = "missing file to review"

	// ErrTooMuchInput is returned when the user provides more than one file.
	ErrTooMuchInput xerrors.Error = "too many files to review; please provide only one file"
)

// ReviewAction is the main action for the application.
func ReviewAction(ctx *cli.Context) error {
	if ctx.Args().Len() < 1 {
		return ErrEmptyInput
	}

	if ctx.Args().Len() > 1 {
		return ErrTooMuchInput
	}

	var (
		key  = ctx.String("key")
		file = ctx.Args().Get(0)
	)

	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	var (
		content = xunsafe.BytesToString(data)
		client  = openai.NewClient(key)
		req     = openai.NewRequest(content)
	)

	resp, err := client.Do(ctx.Context, req)
	if err != nil {
		return fmt.Errorf("failed to get response: %w", err)
	}

	for {
		text, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Fprintf(ctx.App.Writer, "\n")

			break
		}

		if err != nil {
			return fmt.Errorf("failed to get response: %w", err)
		}

		fmt.Fprintf(ctx.App.Writer, "%s", text.Choices[0].Delta.Content)
	}

	return nil
}
