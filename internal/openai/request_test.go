package openai_test

import (
	"testing"

	"github.com/DataDog/documentor/internal/openai"
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
			givePrompt:     openai.MarkdownPrompt,
			expectedPrompt: openai.MarkdownPrompt,
			expectedModel:  "gpt-4o",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := openai.NewRequest("Say 'Hi' and nothing else.", tt.givePrompt, 0.1)

			if req.Model != tt.expectedModel {
				t.Errorf("Got model %v, want %v", req.Model, tt.expectedModel)
			}

			if req.Messages[1].Content != tt.expectedPrompt {
				t.Error("Got a different prompt than expected")
			}
		})
	}
}
