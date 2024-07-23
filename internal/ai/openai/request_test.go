// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

package openai_test

import (
	"testing"

	"github.com/DataDog/documentor/internal/ai/openai"
	"github.com/DataDog/documentor/internal/prompt"
)

func TestNewRequest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		givePrompt     string
		expectedPrompt string
		expectedModel  string
	}{
		{
			name:           "Markdown prompt",
			givePrompt:     prompt.MarkdownPrompt,
			expectedPrompt: prompt.MarkdownPrompt,
			expectedModel:  "gpt-4o-mini",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := openai.NewRequest("Say 'Hi' and nothing else.", "gpt-4o-mini", tt.givePrompt, 0.1)

			if req.Model != tt.expectedModel {
				t.Errorf("Got model %v, want %v", req.Model, tt.expectedModel)
			}

			if req.Messages[1].Content != tt.expectedPrompt {
				t.Error("Got a different prompt than expected")
			}
		})
	}
}

func TestNewRequestWithImage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		giveImage       string
		giveContext     string
		giveFilename    bool
		giveTemperature float32
		expectedModel   string
	}{
		{
			name:            "Image prompt",
			giveImage:       "base64image1",
			giveContext:     "Hello World!",
			giveFilename:    true,
			giveTemperature: 0.1,
			expectedModel:   "gpt-4o-mini",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := openai.NewRequestWithImage(
				tt.giveImage,
				tt.giveContext,
				"gpt-4o-mini",
				"Say 'Hi' and nothing else.",
				tt.giveFilename,
				tt.giveTemperature,
			)

			if req.Model != tt.expectedModel {
				t.Errorf("Got model %v, want %v", req.Model, tt.expectedModel)
			}
		})
	}
}
