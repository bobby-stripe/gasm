//go:build disablew_callstack_overflow_check
// +build disablew_callstack_overflow_check

package buildoptions

const (
	CheckCallStackOverflow = false
	CallStackHeightLimit   = 0
)
