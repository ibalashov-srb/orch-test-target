// Command grep is a minimal grep-style line scanner.
//
// The testable entry point is runGrep; main() is a thin wrapper that parses
// positional arguments via the flag package and routes os.Stdin / os.Stdout
// through to runGrep.
package main

import (
	"bufio"
	"flag"
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

// main parses positional arguments and dispatches to runGrep.
//
// Usage: grep <pattern> [file]
//
// With no arguments, a usage message is written to stderr and the process
// exits with code 2. With one argument, runGrep reads from stdin; with two,
// it reads from the named file. The process exit code is whatever runGrep
// returns (0 = match, 1 = no match, 2 = error).
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage: grep <pattern> [file]")
		os.Exit(2)
	}
	pattern := args[0]
	filePath := ""
	if len(args) >= 2 {
		filePath = args[1]
	}
	os.Exit(runGrep(pattern, filePath, os.Stdin, os.Stdout))
}
