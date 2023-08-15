//go:build tools
// +build tools

package main

import (
	_ "github.com/jstemmer/go-junit-report"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/stringer"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
