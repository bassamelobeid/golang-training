package timer_changes

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Switch go version between 1.22 and 1.23 to see the difference in memory usage
// Behind the scenes, time.After creates a timer that is not freed until it expires.
// And since we are using a large timeout (one hour), the for loop essentially creates a miriad of timers that are not yet freed.
// These timers together use ≈20 MB of memory.

// In Go 1.23, Timers and Tickers that are no longer referred to by the program become eligible for garbage collection immediately,
// even if their Stop methods have not been called.
// So a time.After in a loop will not pile up memory (for long)
func TestTimerMemory(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tokens := make(chan struct{})
	go consumer(ctx, tokens)

	memBefore := getAlloc()

	for i := 0; i < 120000; i++ {
		tokens <- struct{}{}
	}

	memAfter := getAlloc()
	memUsed := memAfter - memBefore
	fmt.Printf("Memory used: %d KB\n", memUsed/1024)

	if memUsed >= 20*1024*1024 {
		t.Errorf("Expected memory usage to be below 20MB, but got %d KB", memUsed/1024)
	}
}

// consumer simulates a consumer that receives tokens
func consumer(ctx context.Context, in <-chan struct{}) {
	for {
		select {
		case <-in:
			// Simulate work
		case <-time.After(time.Hour):
			// Log warning
		case <-ctx.Done():
			return
		}
	}
}

// getAlloc returns the number of bytes of allocated
// heap objects (after garbage collection).
func getAlloc() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}
