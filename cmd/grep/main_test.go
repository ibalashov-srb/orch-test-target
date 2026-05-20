package main

import (
	"os"
	"path/filepath"
	"strings"
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

// TestRunGrep is the table-driven entry point for runGrep behaviour.
// Each sub-test sets up its own input (via tmpFile when file mode is
// exercised) and asserts on runGrep's exit code and captured stdout.
func TestRunGrep(t *testing.T) {
	t.Run("no_match", func(t *testing.T) {
		path := tmpFile(t, "alpha\nbeta\ngamma\n")
		var out strings.Builder
		code := runGrep("zzz", path, nil, &out)
		if code != 1 {
			t.Errorf("no_match: exit code = %d, want 1", code)
		}
		if got := out.String(); got != "" {
			t.Errorf("no_match: stdout = %q, want empty", got)
		}
	})

	t.Run("single_file_match", func(t *testing.T) {
		const matchLine = "beta"
		path := tmpFile(t, "alpha\n"+matchLine+"\ngamma\n")
		var out strings.Builder
		code := runGrep(matchLine, path, nil, &out)
		if code != 0 {
			t.Errorf("single_file_match: exit code = %d, want 0", code)
		}
		want := path + ":2:" + matchLine
		if got := out.String(); !strings.Contains(got, want) {
			t.Errorf("single_file_match: stdout = %q, want it to contain %q", got, want)
		}
	})

	t.Run("missing_file", func(t *testing.T) {
		// Build a path inside a fresh temp dir and never create the file,
		// so os.Open inside runGrep is guaranteed to fail. runGrep must
		// report this with exit code 2 and write nothing to stdout.
		// Stderr output is intentionally not captured or checked.
		path := filepath.Join(t.TempDir(), "does_not_exist.txt")
		var out strings.Builder
		code := runGrep("anything", path, nil, &out)
		if code != 2 {
			t.Errorf("missing_file: exit code = %d, want 2", code)
		}
		if got := out.String(); got != "" {
			t.Errorf("missing_file: stdout = %q, want empty", got)
		}
	})

	t.Run("invalid_regex", func(t *testing.T) {
		// A bad pattern must cause regexp.Compile to fail, which runGrep
		// reports by returning exit code 2. The reader and writer are
		// irrelevant here: runGrep should bail before touching either.
		// Stderr output is intentionally not captured or checked.
		var out strings.Builder
		code := runGrep("[invalid", "", strings.NewReader(""), &out)
		if code != 2 {
			t.Errorf("invalid_regex: exit code = %d, want 2", code)
		}
	})
}
