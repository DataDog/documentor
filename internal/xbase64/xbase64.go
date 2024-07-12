// Package xbase64 extends Go's base64 package.
package xbase64

import (
	"encoding/base64"
)

// EncodeImageToDataURL encodes the given image data as a base64 data URL. It'll
// always assume JPG because the OpenAI API doesn't seem to care.
func EncodeImageToDataURL(data []byte) string {
	base := base64.StdEncoding.EncodeToString(data)

	return "data:image/jpeg;base64," + base
}
