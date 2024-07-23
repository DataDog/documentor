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
	"github.com/DataDog/documentor/internal/errno"
	"github.com/DataDog/documentor/internal/openai"
	"github.com/DataDog/documentor/internal/validate"
	"github.com/DataDog/documentor/internal/xbase64"
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
		context     = ctx.String("context")
		temperature = ctx.Float64("temperature")
		filename    = ctx.Bool("filename")
		file        = ctx.Args().Get(0)
	)

	if !validate.Key(key) {
		return errno.New(errno.ExitUnauthorized, ErrInvalidAPIKey)
	}

	if !validate.Filetype(file, []string{"png", "jpg", "jpeg", "gif"}) {
		return errno.New(errno.ExitInvalidInput, ErrInvalidImageType)
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
		content = xbase64.EncodeImageToDataURL(data)
		client  = openai.NewClient(key)
		req     = openai.NewRequestWithImage(content, context, model, openai.DescribePrompt, filename, float32(temperature))
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
