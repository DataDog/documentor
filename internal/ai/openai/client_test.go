// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

package openai_test

import (
	"context"
	"os"
	"testing"

	"github.com/DataDog/documentor/internal/ai/openai"
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
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var (
				client = openai.NewClient(tt.giveKey)
				req    = openai.NewRequest("hello world", "gpt-4o-mini", "respond with 'Hi' and only that", 0.1)
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
