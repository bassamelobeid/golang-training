package tooling_and_runtime

import (
	"testing"
)

func TestPanicTracebackFormatting(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Recovered from panic: %v", r)
		}
	}()

	panic("what\nhave\nI\ndone")
}
