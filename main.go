// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache-2.0 License. This product includes software developed at
// Datadog (https://www.datadoghq.com/).
// Copyright 2024-Present Datadog, Inc.

package main

import (
	"os"

	"github.com/DataDog/documentor/internal/app"
)

func main() {
	os.Exit(app.Run(os.Args))
}
