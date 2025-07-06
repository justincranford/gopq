# Test Generation Instructions

These instructions guide Copilot and other code generation tools for test code in the gopq project.

## Requirements
- For all Go-specific requirements and best practices, see `go.instructions.md`.
- Cover all exported functions, positive and negative paths, boundary conditions, corner cases, and edge cases
- Include tests for invalid input, tampering, and error propagation
- Ensure coverage of both functional and non-functional requirements, including:
  - Customer-facing and user-facing APIs
  - Deployment, configuration, management, and decommissioning processes
  - Disaster recovery and business continuity scenarios
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
- After splitting, if `<mainfile>_test.go` should be empty. If it is empty, delete it. If it is not empty, stop and warn the user to manually review it.
- Ensure all tests are deterministic and reproducible

## Example: Compliant Test Function

For a Go test example, see `go.instructions.md`.

## Best Practices
- For all Go-specific best practices, see `go.instructions.md`.

---

**Use this file to guide all test code generation for the gopq project.**
