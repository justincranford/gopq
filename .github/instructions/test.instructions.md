# Test Generation Instructions

These instructions guide Copilot and other code generation tools for test code in the gopq project.

## Requirements
- Cover all exported functions, positive and negative paths, and edge cases
- Include tests for invalid input, tampering, and error propagation
- **Always prioritize using the `testify/require` assertion library for all assertions and error checks. Only use `t.Error`, `t.Errorf`, or `t.Fatalf` if `testify/require` is not available or not appropriate. Be consistent within each file.**
- Always check and assert errors, and provide clear, descriptive failure messages
- Always check and assert all return values (and their contents if they are structs, maps, slices, arrays, etc), and provide clear, descriptive failure messages
- Use `TestMain` and `init` for global setup/teardown if needed
- Name all test, fuzz, and benchmark functions clearly and descriptively
- Ensure all tests are deterministic and reproducible
- Unit tests must be in `<mainfile>_unit_test.go`, fuzz tests in `<mainfile>_fuzz_test.go`, and benchmarks in `<mainfile>_bench_test.go`. Remove them from `<mainfile>_test.go`
- After splitting, `<mainfile>_test.go` should be empty or deleted

## Best Practices
- Use Go 1.24.4 features and idioms
- Reference official NIST/FIPS documentation in comments where appropriate
- All code must pass `golangci-lint run --fix` and `gofumpt -l -w .` before commit

---

**Use this file to guide all test code generation for the gopq project.**
