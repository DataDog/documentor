package xbase64_test

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"testing"

	"github.com/DataDog/documentor/internal/xbase64"
)

func TestEncodeImageToDataURL(t *testing.T) {
	t.Parallel()

	testImagePath := filepath.Join("testdata", "dot.gif")

	data, err := os.ReadFile(testImagePath)
	if err != nil {
		t.Fatalf("failed to read test image: %v", err)
	}

	// Manually encode the image to base64 for a valid test case.
	var (
		encodedImage    = base64.StdEncoding.EncodeToString(data)
		expectedDataURL = "data:image/jpeg;base64," + encodedImage
	)

	tests := []struct {
		name string
		give []byte
		want string
	}{
		{
			name: "Valid Image Data",
			give: data,
			want: expectedDataURL,
		},
		{
			name: "Empty Data",
			give: []byte{},
			want: "data:image/jpeg;base64,",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := xbase64.EncodeImageToDataURL(tt.give)
			if got != tt.want {
				t.Fatalf("EncodeImageToDataURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
