package main

import (
	"os"
	"path/filepath"
	"testing"
)

// tmpFile creates a temporary file containing the given content and returns
// its path. The file is removed automatically when the test (and its
// subtests) finish via t.Cleanup.
func tmpFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "input.txt")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("tmpFile: write %s: %v", path, err)
	}
	t.Cleanup(func() {
		// t.TempDir() already removes the directory at test end, but we
		// also explicitly remove the file to satisfy the helper's
		// documented cleanup contract. Ignore "not exist" errors.
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			t.Errorf("tmpFile cleanup: remove %s: %v", path, err)
		}
	})
	return path
}
