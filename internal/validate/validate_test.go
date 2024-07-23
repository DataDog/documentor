// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

package validate_test

import (
	"testing"

	"github.com/DataDog/documentor/internal/validate"
)

func TestFiletype(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		giveFile      string
		giveFiletypes []string
		want          bool
	}{
		{
			name:          "valid file type",
			giveFile:      "test.jpg",
			giveFiletypes: []string{"jpg", "jpeg"},
			want:          true,
		},
		{
			name:          "invalid file type",
			giveFile:      "test.jpg",
			giveFiletypes: []string{"png", "gif"},
			want:          false,
		},
		{
			name:          "no file type",
			giveFile:      "test",
			giveFiletypes: []string{"jpg", "jpeg"},
			want:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := validate.Filetype(tt.giveFile, tt.giveFiletypes)
			if got != tt.want {
				t.Fatalf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		giveKey string
		want    bool
	}{
		{
			name:    "valid key",
			giveKey: "sk-123456789123456789123456789123456789123456789123",
			want:    true,
		},
		{
			name:    "invalid key with invalid prefix",
			giveKey: "ak-123456789123456789123456789123456789123456789123",
			want:    false,
		},
		{
			name:    "invalid key with invalid length",
			giveKey: "sk-123456789",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := validate.Key(tt.giveKey)
			if got != tt.want {
				t.Fatalf("want %v, got %v", tt.want, got)
			}
		})
	}
}
