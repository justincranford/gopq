# Benchmark Instructions

These instructions define standards for writing and maintaining benchmarks in Go projects.

## Requirements
- Place all benchmark tests in `<mainfile>_bench_test.go`.
- Use modern Go idioms, including `b.Loop()` or `b.RunParallel()` for concurrency.
- Always validate and assert all return values in benchmarks.
- Use descriptive variable names and document the purpose of each benchmark.
- Use `testify/require` for assertions in benchmarks.
- All code must pass `golangci-lint run --fix` and `gofumpt -l -w .` before commit.

## Best Practices
- Benchmark both typical and edge-case scenarios.
- Avoid benchmarking code that does I/O or network calls unless necessary.
- Document any non-determinism or variability in results.

---

**Use this file to guide all benchmark code generation and review.**
