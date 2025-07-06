
# Go Language Instructions

Copilot must always follow all requirements and best practices in all other instruction files in this directory (e.g., project, coding, test, bench, fuzz, doc, security). If there is a conflict, project-specific instructions take precedence, followed by domain-specific, then general Go standards.

These instructions consolidate all Go-specific standards and best practices for this and other Go projects. Use this as the single source of truth for Go code, test, benchmark, documentation, and security requirements.

---

## Requirements

- Use Go 1.24.4 or later features and idioms.
- Always check for the latest Go, linter, and dependency versions before starting new work.
- All code must pass `golangci-lint run --fix` and `gofumpt -l -w .` before commit.
- Use `go mod tidy` to manage dependencies and keep `go.mod`/`go.sum` clean.
- Organize code into `cmd/`, `pkg/`, `internal/`, `test/`, and `docs/` directories as appropriate.
- All errors must be wrapped with `fmt.Errorf(..., %w, ...)` for traceability.
- All exported functions, types, and packages must have GoDoc comments.
- Provide usage examples in GoDoc, documentation, and test files.
- Reference official documentation and specifications (e.g., NIST, FIPS) in comments where appropriate.
- Ensure all exported functions are covered by positive and negative tests, fuzz tests, and benchmarks.
- **Always prioritize using the `testify/require` assertion library for all assertions and error checks. Only use `t.Error`, `t.Errorf`, or `t.Fatalf` if `testify/require` is not available or not appropriate. Be consistent within each file.**
- Use `TestMain` and `init` for global setup/teardown if needed
- Name all test, fuzz, and benchmark functions clearly and descriptively
- Unit tests must be in `<mainfile>_unit_test.go`
- Integration tests must be in `<mainfile>_integ_test.go`
- Fuzz tests must be in `<mainfile>_fuzz_test.go`
- Benchmark tests must be in `<mainfile>_bench_test.go`
- If any tests are in `<mainfile>_test.go`, instead of `<mainfile>_unit_test.go`, `<mainfile>_integ_test.go`, `<mainfile>_fuzz_test.go`, or `<mainfile>_bench_test.go`, split them into the appropriate files:
  - Move unit tests to `<mainfile>_unit_test.go`
  - Move integration tests to `<mainfile>_integ_test.go`
  - Move fuzz tests to `<mainfile>_fuzz_test.go`
  - Move benchmark tests to `<mainfile>_bench_test.go`
- After splitting, `<mainfile>_test.go` should be empty. If it is empty, delete it. If it is not empty, stop and warn the user to manually review it.

---

## Best Practices

- Organize code into logical packages and use clear file/folder structure.
- Use dependency injection and interfaces in code where appropriate, to maximize flexibility and testability.
- Use descriptive variable and function names throughout all code and tests. Avoid ambiguous or single-letter names except for idiomatic cases (e.g., `err`).
- Use idiomatic Go error handling and avoid panics except for unrecoverable errors.
- Use `context.Context` for cancellation and deadlines in APIs.
- Use slices and maps idiomatically.
- Prefer composition over inheritance.
- Avoid global state except for constants and configuration.
- Follow the principle of least privilege for all operations. Document all security assumptions and limitations. Review all code for side-channel and timing attack risks.
- Always validate and assert all return values in all tests, errors first and other return values after the error validation. Do not ignore return values with `_`; check the correctness and contents of return values as appropriate.
- Use modern Go idioms, including `b.Loop()` or `b.RunParallel()` for benchmarks instead of legacy for-loops with `b.N`.
- Ensure all unit, integration, and benchmark tests are deterministic and reproducible; only fuzz tests can be a mix of deterministic and/or non-deterministic.
- Avoid benchmarking code that does I/O or network calls unless necessary. Document any non-determinism or variability in results.

---

## Commit Workflow

- **ALWAYS** run `golangci-lint run --fix` and `gofumpt -l -w .` before every commit, and fix all warnings and errors before continuing with the commit, as part of a single chained command. Use `golangci-lint run --fix` to automatically fix all fixable lint issues.
- The canonical commit command is:
  ```
  golangci-lint run --fix; gofumpt -l -w .; git add -A; git commit -m "..."; git push
  ```

---

**Use this file to guide all Go language code generation, testing, benchmarking, documentation, and security review.**
