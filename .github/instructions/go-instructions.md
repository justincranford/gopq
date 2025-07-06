
# Go Language Instructions

Copilot must always follow all requirements and best practices in all other instruction files in this directory (e.g., project, coding, test, bench, fuzz, doc, security). If there is a conflict, project-specific instructions take precedence, followed by domain-specific, then general Go standards.

These instructions consolidate all Go-specific standards and best practices for this and other Go projects. Use this as the single source of truth for Go code, test, benchmark, documentation, and security requirements.

---

## Requirements

- Always use Go 1.24.4 or later features and idioms.
- Always check for the latest Go, linter, and dependency versions before starting new work.
- All code must pass `golangci-lint run --fix` and `gofumpt -l -w .` before commit.
- Use `go mod tidy` to manage dependencies and keep `go.mod`/`go.sum` clean.
- Organize code into `cmd/`, `pkg/`, `internal/`, `test/`, and `docs/` directories as appropriate. For libraries, all public APIs and their tests must be in a public package (e.g., `pq/`), not under `internal/`. Only implementation details or helpers not meant for users should be in `internal/`.
- Always wrap errors with `fmt.Errorf(..., %w, ...)` for traceability.
- Always have GoDoc comments for exported functions, types, and packages.
- Always provide usage examples in GoDoc, documentation, and test files.
- Always cover exported functions with positive and negative unit tests, fuzz tests, and benchmarks; also cover with positive and negative integration tests as appropriate.
- **Always prioritize using the `testify/require` assertion library for all assertions and error checks. Only use `t.Error`, `t.Errorf`, or `t.Fatalf` if `testify/require` is not available or not appropriate. Be consistent within each file.**
- Always name all functions clearly and descriptively in unit tests, integration tests, fuzz tests, and benchmark tests.
- Always put unit tests in `<mainfile>_unit_test.go`
- Always put integration tests in `<mainfile>_integ_test.go`
- Always put fuzz tests in `<mainfile>_fuzz_test.go`
- Always put benchmark tests in `<mainfile>_bench_test.go`
- If any tests are in `<mainfile>_test.go`, instead of `<mainfile>_unit_test.go`, `<mainfile>_integ_test.go`, `<mainfile>_fuzz_test.go`, or `<mainfile>_bench_test.go`, split them into the appropriate files:
  - Move unit tests to `<mainfile>_unit_test.go`
  - Move integration tests to `<mainfile>_integ_test.go`
  - Move fuzz tests to `<mainfile>_fuzz_test.go`
  - Move benchmark tests to `<mainfile>_bench_test.go`
- After splitting, `<mainfile>_test.go` should be empty. If it is empty, delete it. If it is not empty, stop and warn the user to manually review it.
- Always validate and assert all return values in all tests, errors first and then other return values. Do not use `_` to ignore return values in tests; check the correctness and contents of return values as appropriate.
- Follow the principle of least privilege for all operations. Document all security assumptions and limitations. Review all code for side-channel and timing attack risks.
- Use modern Go idioms, including `b.Loop()` or `b.RunParallel()` for benchmarks instead of legacy for-loops with `b.N`.

---

## Best Practices

- Use dependency injection and interfaces in code where appropriate, to maximize flexibility and testability.
- Organize code into logical packages and use clear file/folder structure.
- Use descriptive variable and function names throughout all code and tests. Avoid ambiguous or single-letter names except for idiomatic cases (e.g., `err`).
- Use idiomatic Go error handling and avoid panics except for unrecoverable errors.
- Use `context.Context` for cancellation and deadlines in APIs.
- Use `TestMain` and `init` for global setup/teardown if needed
- Use slices and maps idiomatically.
- Avoid global state except for constants and configuration.
- Ensure all unit, integration, and benchmark tests are deterministic and reproducible; only fuzz tests can be a mix of deterministic and/or non-deterministic.
- Avoid benchmarking code that does I/O or network calls unless necessary. Document any non-determinism or variability in results.
- Prefer composition over inheritance.

---

## Commit Workflow

- **ALWAYS** run `golangci-lint run --fix` and `gofumpt -l -w .` before every commit, and fix all warnings and errors before continuing with the commit, as part of a single chained command. Use `golangci-lint run --fix` to automatically fix all fixable lint issues.
- The canonical commit command is:
  ```
  golangci-lint run --fix; gofumpt -l -w .; git add -A; git commit -m "..."; git push
  ```

---

**Use this file to guide all Go language code generation, testing, benchmarking, documentation, and security review.**
