# Test Generation Instructions

These instructions guide Copilot and other code generation tools for test code in the gopq project.

## Requirements
- Use clear, descriptive variable names throughout all test code. Avoid ambiguous or single-letter names except for idiomatic cases (e.g., err).
- Always validate and assert all return values from functions under test. Do not ignore return values with `_`; check their correctness and contents as appropriate.
- Use modern Go idioms and APIs, including `b.Loop()` for benchmarks instead of legacy for-loops with `b.N`.
- Cover all exported functions, positive and negative paths, and edge cases
- Include tests for invalid input, tampering, and error propagation
- **Always prioritize using the `testify/require` assertion library for all assertions and error checks. Only use `t.Error`, `t.Errorf`, or `t.Fatalf` if `testify/require` is not available or not appropriate. Be consistent within each file.**
- Always check and assert errors, and provide clear, descriptive failure messages.
- Always check and assert all return values (and their contents if they are structs, maps, slices, arrays, etc), and provide clear, descriptive failure messages. Do not ignore any return value.
- Use `TestMain` and `init` for global setup/teardown if needed
- Ensure all tests are deterministic and reproducible
- Name all test, fuzz, and benchmark functions clearly and descriptively
- Unit tests must be in `<mainfile>_unit_test.go`
- Fuzz tests must be in `<mainfile>_fuzz_test.go`
- Benchmark tests must be in `<mainfile>_bench_test.go`
- If any tests are in `<mainfile>_test.go`, instead of `<mainfile>_unit_test.go`, `<mainfile>_fuzz_test.go`, or `<mainfile>_bench_test.go`, split them into the appropriate files:
  - Move unit tests to `<mainfile>_unit_test.go`
  - Move fuzz tests to `<mainfile>_fuzz_test.go`
  - Move benchmark tests to `<mainfile>_bench_test.go`
- After splitting, if `<mainfile>_test.go` should be empty. If it is empty, delete it. If it is not empty, stop and warn the user to manually review it.

## Example: Compliant Test Function

Here is an example of a compliant unit test using `testify/require`:

```go
import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestExample_Addition_Unit(t *testing.T) {
    result := 2 + 2
    require.Equal(t, 4, result, "Addition result should be 4")
}
```

## Best Practices
- Use Go 1.24.4 features and idioms
- Reference official NIST/FIPS documentation in comments where appropriate
- All code must pass `golangci-lint run --fix` and `gofumpt -l -w .` before commit
- All exported functions must have both positive and negative test cases.

---

**Use this file to guide all test code generation for the gopq project.**
