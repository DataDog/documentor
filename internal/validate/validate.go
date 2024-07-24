// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

// Package validate contains functions and utilities for validating input data.
package validate

import (
	"path/filepath"
	"strings"
)

// Filetype validates that the given file is of the correct type.
func Filetype(file string, allowedTypes []string) bool {
	filetype := filepath.Ext(file)
	filetype = strings.ToLower(filetype)

	if filetype != "" {
		filetype = filetype[1:]
	}

	for _, allowed := range allowedTypes {
		if filetype == allowed {
			return true
		}
	}

	return false
}

// Key validates that the given API key is of the correct format.
func Key(key string) bool {
	if len(key) < 51 {
		return false
	}

	if !strings.HasPrefix(key, "sk-") {
		return false
	}

	return true
}
