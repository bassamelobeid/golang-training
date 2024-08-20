package os_operations

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFS(t *testing.T) {
	// Create a temporary source directory
	srcDir, err := os.MkdirTemp("", "src")
	if err != nil {
		t.Fatalf("Failed to create temp source directory: %v", err)
	}
	defer os.RemoveAll(srcDir)

	// Create a test file in the source directory
	testFilePath := filepath.Join(srcDir, "test.txt")
	err = os.WriteFile(testFilePath, []byte("Hello, World!"), os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Define destination directory
	dstDir := filepath.Join(os.TempDir(), "dst")

	// Copy the directory
	// FYI: this won't work for symlinks
	err = os.CopyFS(dstDir, os.DirFS(srcDir))
	if err != nil {
		t.Fatalf("Failed to copy directory: %v", err)
	}

	// Verify the file was copied
	copiedFilePath := filepath.Join(dstDir, "test.txt")
	if _, err := os.Stat(copiedFilePath); os.IsNotExist(err) {
		t.Errorf("Expected file %s to be copied, but it does not exist", copiedFilePath)
	}
}
