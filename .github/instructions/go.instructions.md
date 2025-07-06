# Go Language Instructions

These instructions consolidate all Go-specific standards and best practices for this and other Go projects. Use this as the single source of truth for Go code, test, benchmark, documentation, and security requirements.

- Use modern Go idioms, including `b.Loop()` or `b.RunParallel()` for benchmarks instead of legacy for-loops with `b.N`.
- Always validate and assert all return values in benchmarks.
- Use descriptive variable names and document the purpose of each benchmark.
- Use `testify/require` for assertions in all test, fuzz, and benchmark code unless otherwise specified. Only use `t.Error`, `t.Errorf`, or `t.Fatalf` if `testify/require` is not available or not appropriate. Be consistent within each file.
- Use Go 1.24.4 or later features and idioms.
- Use clear, descriptive variable and function names throughout all code and tests. Avoid ambiguous or single-letter names except for idiomatic cases (e.g., err).
- All errors must be wrapped with `fmt.Errorf(..., %w, ...)` for traceability.
- All exported functions, types, and packages must have GoDoc comments.
- Provide usage examples in GoDoc, documentation, and test files.
- All code must pass `golangci-lint run --fix` and `gofumpt -l -w .` before commit.
- Use `go mod tidy` to manage dependencies and keep `go.mod`/`go.sum` clean.
- Reference official documentation and specifications (e.g., NIST, FIPS) in comments where appropriate.
- Ensure all exported functions are covered by positive and negative tests, fuzz tests, and benchmarks.
- Use `testify/require` for assertions in all test, fuzz, and benchmark code unless otherwise specified. Only use `t.Error`, `t.Errorf`, or `t.Fatalf` if `testify/require` is not available or not appropriate. Be consistent within each file.
- Always validate and assert all return values from functions under test. Do not ignore return values with `_`; check their correctness and contents as appropriate.
- Use modern Go idioms and APIs, including `b.Loop()` or `b.RunParallel()` for benchmarks instead of legacy for-loops with `b.N`.
- Place all benchmark tests in `<mainfile>_bench_test.go`.
- Unit tests must be in `<mainfile>_unit_test.go`, fuzz tests in `<mainfile>_fuzz_test.go`, and benchmarks in `<mainfile>_bench_test.go`. After splitting, if `<mainfile>_test.go` is empty, delete it. If not, stop and warn the user to manually review it.
- Ensure all tests and benchmarks are deterministic and reproducible. Benchmark both typical and edge-case scenarios.

- Avoid benchmarking code that does I/O or network calls unless necessary. Document any non-determinism or variability in results.
- Never ignore returned errors or values; always check and assert them.
- Use idiomatic Go error handling and avoid panics except for unrecoverable errors.
- Organize code into logical packages and use clear file/folder structure.
- Use dependency injection and interfaces for testability where appropriate.
- Prefer composition over inheritance.
- Use slices and maps idiomatically.
- Avoid global state except for constants and configuration.
- Use context.Context for cancellation and deadlines in APIs.
- Avoid benchmarking code that does I/O or network calls unless necessary. Document any non-determinism or variability in results.
- Follow the principle of least privilege for all operations. Document all security assumptions and limitations. Review all code for side-channel and timing attack risks.

---

**Use this file to guide all Go language code generation, testing, benchmarking, documentation, and security review.**
