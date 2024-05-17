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
	"github.com/DataDog/documentor/internal/errno"
	"github.com/DataDog/documentor/internal/openai"
	"github.com/urfave/cli/v2"
)

const (
	// ErrEmptyInput is returned when no file is provided by the user.
	ErrEmptyInput xerrors.Error = "missing file to review"

	// ErrEmptyAPIKey is returned when the user does not provide an OpenAI API
	// key.
	ErrEmptyAPIKey xerrors.Error = "missing OpenAI API key"

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
		key  = ctx.String("key")
		file = ctx.Args().Get(0)
	)

	if key == "" {
		return errno.New(errno.ExitUnauthorized, ErrEmptyAPIKey)
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
		content = xunsafe.BytesToString(data)
		client  = openai.NewClient(key)
		req     = openai.NewRequest(content)
	)

	resp, err := client.Do(ctx.Context, req)
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return errno.New(errno.ExitTimeout, err)
		}

		return errno.New(errno.ExitAPIError, fmt.Errorf("failed to get response: %w", err))
	}

	for {
		text, err := resp.Recv()
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
