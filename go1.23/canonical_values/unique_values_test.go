package unique_values

import (
	"runtime"
	"testing"
	"unique"
)

func TestMemoryWithoutCanonicalization(t *testing.T) {
	const nWords = 10000
	const nDistinct = 100
	const wordLen = 4

	generate := wordGen(nDistinct, wordLen)
	memBefore := getAlloc()

	words := make([]string, nWords)
	for i := range words {
		words[i] = generate()
	}

	memAfter := getAlloc()
	memUsed := memAfter - memBefore
	t.Logf("Memory used without canonicalization: %d KB", memUsed/1024)

	if memUsed <= 100*1024 {
		t.Errorf("Expected memory usage to exceed 100 KB, but got %d KB", memUsed/1024)
	}
}

func TestMemoryWithCanonicalization(t *testing.T) {
	const nWords = 10000
	const nDistinct = 100
	const wordLen = 4

	generate := wordGen(nDistinct, wordLen)
	memBefore := getAlloc()

	words := make([]unique.Handle[string], nWords)
	for i := range words {
		words[i] = unique.Make(generate())
	}

	memAfter := getAlloc()
	memUsed := memAfter - memBefore
	t.Logf("Memory used with canonicalization: %d KB", memUsed/1024)

	if memUsed >= 600*1024 {
		t.Errorf("Expected memory usage below 600 KB, but got %d KB", memUsed/1024)
	}
}

// Helper functions

// wordGen returns a generator that produces deterministic words from a fixed character set.
func wordGen(nDistinct int, wordLen int) func() string {
	vocab := make([]string, nDistinct)
	charSet := "abcd" // A simple fixed character set
	for i := 0; i < nDistinct; i++ {
		vocab[i] = generateWord(i, wordLen, charSet)
	}
	index := 0
	return func() string {
		word := vocab[index%len(vocab)]
		index++
		return word
	}
}

// generateWord creates a word from a fixed character set using an index.
func generateWord(index, wordLen int, charSet string) string {
	word := make([]byte, wordLen)
	charSetLen := len(charSet)
	for i := 0; i < wordLen; i++ {
		word[i] = charSet[index%charSetLen]
		index /= charSetLen
	}
	return string(word)
}

// getAlloc is a placeholder to simulate memory allocation measurement.
func getAlloc() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}
