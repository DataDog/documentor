// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

// Package app is the main package for the application.
package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/DataDog/documentor/internal/ai"
	"github.com/DataDog/documentor/internal/ai/openai"
	"github.com/DataDog/documentor/internal/errno"
	"github.com/DataDog/documentor/internal/meta"
	"github.com/urfave/cli/v2"
)

// Run is the entry point for the application.
func Run(args []string) int {
	app := cli.NewApp()
	app.Name = meta.Name
	app.Version = meta.Version
	app.Usage = meta.Description
	app.HideHelpCommand = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "key",
			Aliases: []string{"k"},
			Usage:   "the API key to use",
			EnvVars: []string{
				"DOCUMENTOR_KEY",
			},
		},
		&cli.StringFlag{
			Name:    "provider",
			Aliases: []string{"p"},
			Usage:   "the AI provider to use",
			Value:   ai.ProviderOpenAI,
			EnvVars: []string{
				"DOCUMENTOR_PROVIDER",
			},
		},
		&cli.StringFlag{
			Name:    "model",
			Aliases: []string{"m"},
			Usage:   "the AI model to use",
			Value:   openai.DefaultModel,
			EnvVars: []string{
				"DOCUMENTOR_MODEL",
			},
		},
		&cli.Float64Flag{
			Name:    "temperature",
			Aliases: []string{"t"},
			Usage:   "the temperature to use for the model",
			Value:   0.8,
			EnvVars: []string{
				"DOCUMENTOR_TEMPERATURE",
			},
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "review",
			Aliases: []string{"r"},
			Usage:   "review technical documentation",
			Action:  ReviewAction,
		},
		{
			Name:    "describe",
			Aliases: []string{"d"},
			Usage:   "describe an image and generate alt text",
			Action:  DescribeAction,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "context",
					Aliases: []string{"c"},
					Usage:   "the context to use for the image",
				},
				&cli.BoolFlag{
					Name:    "filename",
					Aliases: []string{"f"},
					Usage:   "whether to generate a filename",
					Value:   false,
				},
			},
		},
		{
			Name:    "draft",
			Aliases: []string{"D"},
			Usage:   "draft new documentation based on the provided notes",
			Action:  DraftAction,
		},
	}

	if err := app.Run(args); err != nil {
		var exitErr *errno.Error

		if errors.As(err, &exitErr) {
			fmt.Fprintf(os.Stderr, "error: %q\n", err)

			return exitErr.Code()
		}

		fmt.Fprintf(os.Stderr, "error: unknown error: %q\n", err)

		return int(errno.ExitError)
	}

	return 0
}
