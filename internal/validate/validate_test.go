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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := validate.Filetype(tt.giveFile, tt.giveFiletypes)
			if got != tt.want {
				t.Fatalf("want %v, got %v", tt.want, got)
			}
		})
	}
}
