# Contributing to orch-test-target

Welcome, and thanks for your interest in contributing to **orch-test-target**,
a small Go 1.22 `grep`-like command-line utility used as a smoke-test target
for the orchestrator pipeline. The notes below describe how to set up a
development environment, the conventions we follow, and how to get changes
merged.

## Prerequisites

- Go 1.22 or newer (the module declares `go 1.22` in `go.mod`).
- `git` for cloning the repository and preparing pull requests.
- A POSIX-like shell is recommended for running the commands below.

## Getting Started

1. Fork the repository on GitHub, then clone your fork:
   `git clone https://github.com/<you>/orch-test-target.git`
2. Add the upstream remote so you can stay current with `main`:
   `git remote add upstream https://github.com/ibalashov-srb/orch-test-target.git`
3. From the repo root, build and test everything:
   - `go build ./...`
   - `go test ./...`

## Code Style

- Format every Go file before committing: `go fmt ./...`.
- Run `go vet ./...` and resolve any reported issues.
- Keep changes minimal and focused; avoid drive-by reformatting of code
  unrelated to your patch.

## Pull-Request Workflow

- Branch from `main` (e.g. `git checkout -b fix/usage-error main`).
- Target the upstream `main` branch when opening the PR.
- Keep PRs small and self-contained — one logical change per PR.
- In the PR description, explain *what* the change does and *why* it is
  needed; link any relevant issues.
- Make sure `go build ./...` and `go test ./...` pass locally before
  requesting review.

## Commit Messages

- Use the imperative mood in the subject line (e.g. "Add stdin support",
  not "Added" or "Adds").
- Keep the subject line to 72 characters or fewer; wrap the body at ~72.
- Reference related issues using `#N` (for example, `Fixes #42`) when
  applicable.

## Reporting Issues

- Search existing issues first to avoid filing duplicates.
- Include your OS (and version) and the output of `go version` so we know
  which Go 1.22+ toolchain you used.
- Provide clear, minimal reproduction steps, the command you ran, the
  output you saw, and the output you expected.
