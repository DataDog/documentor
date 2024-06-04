package openai_test

import (
	"context"
	"os"
	"testing"

	"github.com/DataDog/documentor/internal/openai"
)

func TestClient_Do(t *testing.T) {
	t.Parallel()

	key, found := os.LookupEnv("DOCUMENTOR_KEY")
	if !found || key == "" {
		t.Skip("missing required environment variable: DOCUMENTOR_KEY")
	}

	tests := []struct {
		name    string
		giveKey string
		wantErr bool
	}{
		{
			name:    "Valid request",
			giveKey: key,
			wantErr: false,
		},
		{
			name:    "Invalid request",
			giveKey: "invalid-api-key",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var (
				client = openai.NewClient(tt.giveKey)
				req    = openai.NewRequest("hello world", "respond with 'Hi' and only that")
			)

			_, err := client.Do(context.Background(), req)
			if tt.wantErr && err == nil {
				t.Error("Expected error but got none")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("Didn't expect errors, got %v", err)
			}
		})
	}
}
