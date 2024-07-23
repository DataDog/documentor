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
