# Go Language Instructions

These instructions define Go-specific standards and best practices for all Go code in this project and others.

## Requirements
- Use Go 1.24.4 or later features and idioms.
- Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
- Use `go mod tidy` to manage dependencies and keep `go.mod`/`go.sum` clean.
- Use `golangci-lint run --fix` and `gofumpt -l -w .` before every commit.
- Use `testify/require` for assertions in tests unless otherwise specified.
- All exported functions and types must have GoDoc comments.

## Best Practices
- Prefer composition over inheritance.
- Use slices and maps idiomatically.
- Avoid global state except for constants and configuration.
- Use context.Context for cancellation and deadlines in APIs.

---

**Use this file to guide all Go language code generation and review.**
