# Fuzz Test Instructions

These instructions define standards for writing and maintaining fuzz tests.

## Requirements
- For all Go-specific requirements and best practices, see `go-instructions.md`.
- Fuzz tests must validate input handling, robustness, and reliability of all exported functions and APIs.
- Cover edge cases, malformed input, and tampering scenarios.
- Ensure fuzz tests exercise recovery from invalid or unexpected input.

## Best Practices
- Cover aspects related to:
  - Input validation and sanitization
  - Robustness against malformed or adversarial input
  - Reliability under repeated or random input
  - Recovery from panics, errors, or resource exhaustion
- Document any limitations or known gaps in fuzz coverage.

---

**Use this file to guide all fuzz test code generation and review.**
