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
