package timer_changes

import (
	"fmt"
	"testing"
	"time"
)

// Switch go version between 1.22 and 1.23 to see the difference in timer behavior
func TestTimerResetBehavior(t *testing.T) {
	const timeout = 10 * time.Millisecond
	tm := time.NewTimer(timeout)
	time.Sleep(20 * time.Millisecond)

	start := time.Now()
	tm.Reset(timeout)
	<-tm.C

	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Time elapsed: %dms\n", elapsed)

	// Expecting 10ms after the reset due to the fix in Go 1.23
	if elapsed < 10 {
		t.Errorf("Expected elapsed time >= 10ms, but got %dms", elapsed)
	}
}
