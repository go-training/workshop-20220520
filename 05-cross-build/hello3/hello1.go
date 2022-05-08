//go:build !go1.18
// +build !go1.18

package hello3

import (
	"runtime"
)

// Hello is hello world
func Hello() string {
	return "build on " + runtime.Version()
}
