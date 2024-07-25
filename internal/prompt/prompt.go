// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

// Package prompt holds prompts to be used by AI.
package prompt

import (
	_ "embed"
)

//go:embed data/review-prompt.txt
var MarkdownPrompt string

//go:embed data/describe-prompt.txt
var DescribePrompt string

//go:embed data/draft-prompt.txt
var DraftPrompt string
