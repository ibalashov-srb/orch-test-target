// Command grep is a minimal grep-style line scanner.
//
// This file currently provides the testable entry point runGrep. A thin
// main() wrapper that parses positional arguments (and routes os.Stdin /
// os.Stdout through to runGrep) is added in a subsequent change.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

// runGrep is the testable entry point for the grep CLI.
//
//	pattern  – Go regular-expression string.
//	filePath – path to the file to scan. The empty string or "-" selects
//	           stdin mode and the provided `in` reader is used instead.
//	in       – reader consulted in stdin mode (ignored when filePath names
//	           a real file).
//	out      – writer that receives matching lines.
//
// Matching lines in file mode are written as "<filename>:<lineno>:<line>\n"
// with 1-based line numbers; in stdin mode they are written as
// "<lineno>:<line>\n" with no filename prefix.
//
// The return value is the process exit code:
//
//	0 – at least one line matched.
//	1 – input was read successfully but no line matched.
//	2 – pattern failed to compile, or the named file could not be opened.
//
// On a code-2 error an explanatory message is written to os.Stderr.
func runGrep(pattern, filePath string, in io.Reader, out io.Writer) int {
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "grep: invalid pattern: %v\n", err)
		return 2
	}

	var reader io.Reader
	stdinMode := filePath == "" || filePath == "-"
	if stdinMode {
		reader = in
	} else {
		f, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "grep: %v\n", err)
			return 2
		}
		defer f.Close()
		reader = f
	}

	matched := 0
	scanner := bufio.NewScanner(reader)
	lineno := 0
	for scanner.Scan() {
		lineno++
		line := scanner.Text()
		if !re.MatchString(line) {
			continue
		}
		matched++
		if stdinMode {
			fmt.Fprintf(out, "%d:%s\n", lineno, line)
		} else {
			fmt.Fprintf(out, "%s:%d:%s\n", filePath, lineno, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "grep: read error: %v\n", err)
		return 2
	}

	if matched == 0 {
		return 1
	}
	return 0
}

// main is a minimal placeholder so this file compiles as package main.
// A full argument-parsing implementation is added in a follow-up bead;
// adding it here would require importing `flag`, which is outside this
// bead's stdlib-import allowlist.
func main() {
	os.Exit(runGrep("", "", os.Stdin, os.Stdout))
}
