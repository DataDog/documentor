// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

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
