# General Coding Instructions

These instructions provide general, reusable coding standards for all Go projects using Copilot or other code generation tools.

## Requirements
- Use clear, descriptive variable and function names throughout all code.
- Use Go 1.24.4 features and idioms.
- All errors must be wrapped with `fmt.Errorf(..., %w, ...)` for traceability.
- Add GoDoc comments to all exported functions, types, and packages.
- Provide usage examples in documentation and tests.
- All code must pass `golangci-lint run --fix` and `gofumpt -l -w .` before commit.
- Reference official documentation in comments where appropriate.
- Ensure all exported functions are covered by positive and negative tests, fuzz tests, and benchmarks.

## Best Practices
- Never ignore returned errors or values; always check and assert them.
- Use idiomatic Go error handling and avoid panics except for unrecoverable errors.
- Organize code into logical packages and use clear file/folder structure.
- Use dependency injection and interfaces for testability where appropriate.

---

**Use this file to guide all general code generation for Go projects.**
