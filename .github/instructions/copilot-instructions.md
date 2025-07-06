# Copilot/AI Code Generation Instructions

These instructions define how Copilot and other AI code generation tools should operate in this project.

## Requirements
- Always follow all project-specific and general instructions files (e.g., coding, go, test, bench, doc, security).
- Use PowerShell/Windows command syntax for all git and code quality commands.
- Always run `golangci-lint run --fix` and `gofumpt -l -w .` before every commit.
- Use descriptive variable names and wrap all errors with `fmt.Errorf(..., %w, ...)`.
- Prioritize `testify/require` for assertions in all test code.
- Never ignore return values or errors.
- Reference official documentation and specifications in comments where appropriate.

## Best Practices
- Prefer reusable, idiomatic, and maintainable code.
- Organize code and documentation for clarity and discoverability.

---

**Use this file to guide all Copilot/AI code generation in this project.**
