//go:build stress
// +build stress

package tests

import (
	"runtime"
)

func init() {

	defaultTimeout *= 5

	runtime.GOMAXPROCS(runtime.GOMAXPROCS(0) * 2)
}
