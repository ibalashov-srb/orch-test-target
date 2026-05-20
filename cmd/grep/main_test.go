package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// helper: run runGrep and return (exitCode, captured output)
func runGrepCapture(pattern, filePath string, input string) (int, string) {
	var out bytes.Buffer
	var in *strings.Reader
	if input != "" {
		in = strings.NewReader(input)
	} else {
		in = strings.NewReader("")
	}
	code := runGrep(pattern, filePath, in, &out)
	return code, out.String()
}

// ---------------------------------------------------------------------------
// stdin mode (filePath == "" or "-")
// ---------------------------------------------------------------------------

func TestStdinMode_Match(t *testing.T) {
	code, out := runGrepCapture("hello", "", "hello world\ngoodbye\nhello again\n")
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	want := "1:hello world\n3:hello again\n"
	if out != want {
		t.Errorf("expected output %q, got %q", want, out)
	}
}

func TestStdinMode_NoMatch(t *testing.T) {
	code, out := runGrepCapture("xyz", "", "hello world\ngoodbye\n")
	if code != 1 {
		t.Errorf("expected exit code 1, got %d", code)
	}
	if out != "" {
		t.Errorf("expected empty output, got %q", out)
	}
}

func TestStdinMode_DashFilePath(t *testing.T) {
	// filePath == "-" should also trigger stdin mode
	var out bytes.Buffer
	in := strings.NewReader("match me\nskip\nmatch again\n")
	code := runGrep("match", "-", in, &out)
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	want := "1:match me\n3:match again\n"
	if out.String() != want {
		t.Errorf("expected %q, got %q", want, out.String())
	}
}

func TestStdinMode_LineNumbers1Based(t *testing.T) {
	// The first matching line is line 1
	code, out := runGrepCapture("a", "", "a\n")
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	if !strings.HasPrefix(out, "1:") {
		t.Errorf("expected line numbers to be 1-based; got %q", out)
	}
}

// ---------------------------------------------------------------------------
// file mode
// ---------------------------------------------------------------------------

func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "input.txt")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	return path
}

func TestFileMode_Match(t *testing.T) {
	path := writeTempFile(t, "apple\nbanana\napricot\n")
	var out bytes.Buffer
	code := runGrep("^a", path, nil, &out)
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	want := path + ":1:apple\n" + path + ":3:apricot\n"
	if out.String() != want {
		t.Errorf("expected %q, got %q", want, out.String())
	}
}

func TestFileMode_NoMatch(t *testing.T) {
	path := writeTempFile(t, "apple\nbanana\n")
	var out bytes.Buffer
	code := runGrep("xyz", path, nil, &out)
	if code != 1 {
		t.Errorf("expected exit code 1, got %d", code)
	}
	if out.String() != "" {
		t.Errorf("expected empty output, got %q", out.String())
	}
}

func TestFileMode_Format(t *testing.T) {
	// Each matching line must be formatted as "<filename>:<lineno>:<line>"
	path := writeTempFile(t, "line one\nline two\nline three\n")
	var out bytes.Buffer
	code := runGrep("two", path, nil, &out)
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	expected := path + ":2:line two\n"
	if out.String() != expected {
		t.Errorf("expected %q, got %q", expected, out.String())
	}
}

func TestFileMode_FileNotFound(t *testing.T) {
	var out bytes.Buffer
	code := runGrep("pattern", "/nonexistent/path/file.txt", nil, &out)
	if code != 2 {
		t.Errorf("expected exit code 2 for missing file, got %d", code)
	}
}

// ---------------------------------------------------------------------------
// error cases
// ---------------------------------------------------------------------------

func TestInvalidPattern(t *testing.T) {
	code, _ := runGrepCapture("[invalid(", "", "anything\n")
	if code != 2 {
		t.Errorf("expected exit code 2 for invalid regexp, got %d", code)
	}
}

// ---------------------------------------------------------------------------
// edge cases
// ---------------------------------------------------------------------------

func TestEmptyInput(t *testing.T) {
	code, out := runGrepCapture("hello", "", "")
	if code != 1 {
		t.Errorf("expected exit code 1 for empty input, got %d", code)
	}
	if out != "" {
		t.Errorf("expected empty output, got %q", out)
	}
}

func TestMatchesEveryLine(t *testing.T) {
	code, out := runGrepCapture(".", "", "a\nb\nc\n")
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	want := "1:a\n2:b\n3:c\n"
	if out != want {
		t.Errorf("expected %q, got %q", want, out)
	}
}

func TestFileMode_SingleMatchLastLine(t *testing.T) {
	path := writeTempFile(t, "skip\nskip\nfound\n")
	var out bytes.Buffer
	code := runGrep("found", path, nil, &out)
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	expected := path + ":3:found\n"
	if out.String() != expected {
		t.Errorf("expected %q, got %q", expected, out.String())
	}
}
