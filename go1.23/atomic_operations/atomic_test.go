package atomic_operations

import (
	"sync/atomic"
	"testing"
)

func TestAtomicAndOr(t *testing.T) {
	const (
		modeRead  = 0b100
		modeWrite = 0b010
		modeExec  = 0b001
	)

	var mode atomic.Int32
	mode.Store(modeRead)
	old := mode.Or(modeWrite)

	expectedOld := int32(modeRead)
	expectedNew := int32(modeRead | modeWrite)

	if old != expectedOld || mode.Load() != expectedNew {
		t.Errorf("Expected old: %b, new: %b, but got old: %b, new: %b", expectedOld, expectedNew, old, mode.Load())
	}
}
