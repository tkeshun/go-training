//go:build tools
// +build tools

package pkg

// tools.go

import (
	_ "github.com/cweill/gotests/gotests" // gotestsの依存を宣言
)
