# MANDATORY GIT WORKFLOW RULES
- ALWAYS use PowerShell terminal and Windows command syntax for all git and code quality commands, regardless of the user's OS or default shell.
- ALWAYS run `golangci-lint run --fix` and `gofumpt -l -w .` before every commit, and fix all warnings and errors before continuing with the commit, as part of a single chained command. Use `golangci-lint run --fix` to automatically fix all fixable lint issues.
- NEVER use Git Bash or Unix shell syntax for any git or code quality operation.
- The canonical commit command is:
  ```
  golangci-lint run --fix; gofumpt -l -w .; git add -A; git commit -m "..."; git push
  ```
- These rules override all other shell/environment logic.
- ALWAYS use descriptive variable names throughout all code, tests, and documentation. Avoid single-letter or ambiguous names except for idiomatic cases (e.g., error as `err`).

# Copilot Instructions for gopq

## Project Overview
This project demonstrates reusable utility methods for executing post-quantum safe algorithms, specifically focusing on NIST FIPS 140-3 approved algorithms:
- **ML-DSA** (Multivariate Lattice Digital Signature Algorithm)
- **ML-KEM** (Multivariate Lattice Key Encapsulation Mechanism)

The project is implemented in Go, targeting version **1.24.4** and using the latest APIs and dependencies. It includes:
- Test examples
- Fuzz tests
- Performance benchmark tests
- Comprehensive code comments
- End-user documentation

## Coding Standards
- **Error Wrapping:** All returned errors must be wrapped using `fmt.Errorf` with the original error as `%w` for traceability.
- **Go Version:** All code must use Go 1.24.4 features and APIs. Reference the [Go Language Specification](https://go.dev/ref/spec) and [Standard Library](https://pkg.go.dev/std) for up-to-date usage.
- **Dependencies:** Always use the latest stable versions of all dependencies. Update dependencies regularly.
- **Formatting:**
  - Use the latest versions of `golangci-lint` and `gofumpt`.
  - Enable all optional formatters and linters in both tools.
  - Code must pass all enabled linters and formatters before merging.

## Project Structure
- `cmd/` — CLI or main entry points
- `pkg/` — Reusable packages for ML-DSA and ML-KEM utilities
- `internal/` — Internal helpers and utilities
- `test/` — Test examples, fuzz tests, and benchmarks
- `docs/` — End-user documentation

## Testing
- Include unit tests for all exported functions, and always move unit tests to a separate file named `<mainfile>_unit_test.go`. Remove them from `<mainfile>_test.go`.
- Add fuzz tests for cryptographic primitives and edge cases, and always move fuzz tests to a separate file named `<mainfile>_fuzz_test.go`. Remove them from `<mainfile>_test.go`.
- Provide performance benchmarks for all major algorithms, and always move benchmark tests to a separate file named `<mainfile>_bench_test.go`. Remove them from `<mainfile>_test.go`.
- After splitting, `<mainfile>_test.go` should be empty or deleted.
- Ensure tests are reproducible and well-documented.

## Documentation
- All exported functions and types must have GoDoc comments.
- Provide usage examples in documentation and tests.
- Maintain a `README.md` with high-level usage and security notes.


## Contribution Guidelines
- All code must be reviewed and pass CI checks (lint, format, test, fuzz, benchmark).
- Use descriptive commit messages and PR descriptions.
- Reference relevant NIST and FIPS documentation in code comments where appropriate.
- **Before every git commit, always run `golangci-lint run` and `gofumpt -l -w .` and ensure there are no lint or formatting errors.**
- **When requesting a git commit, always use Git Bash terminal and syntax, and run all commands (`golangci-lint run && gofumpt -l -w . && git add -A && git commit -m "..." && git push`) as a single operation, without prompting for each step.**

---

**Note:**
- Always check for the latest Go, linter, and dependency versions before starting new work.
- Follow best practices for cryptographic code and reference official specifications.
