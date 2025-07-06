# General Coding Instructions

These instructions provide general, reusable coding standards for all Go projects.

## Requirements
- Use clear, descriptive variable and function names throughout all code.
- Provide usage examples in documentation and tests.
- Reference official documentation and specifications (e.g., NIST, FIPS) in comments where appropriate.
- Ensure all exported functions are covered by positive and negative tests, fuzz tests, and benchmarks.
- For all Go-specific requirements and best practices, see `go.instructions.md`.

## Best Practices
- Never ignore returned errors or values; always check and validate them.
- Organize code into logical packages and use clear file/folder structure.
- Use dependency injection and interfaces for testability where appropriate.
- For all Go-specific best practices, see `go.instructions.md`.

---

**Use this file to guide all general code generation for Go projects.**
