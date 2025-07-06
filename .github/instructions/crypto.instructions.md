# Crypto Code Generation Instructions

These instructions guide Copilot and other code generation tools for cryptographic code in the gopq project.

## Requirements
- Use only NIST/FIPS-approved post-quantum algorithms (ML-KEM, ML-DSA, Kyber, Dilithium, etc.)
- Use Go 1.24.4 features and idioms
- All errors must be wrapped with `fmt.Errorf(..., %w, ...)`
- Use descriptive variable names throughout
- Reference official NIST/FIPS documentation in comments where appropriate
- All code must pass `golangci-lint run --fix` and `gofumpt -l -w .` before commit
- Add GoDoc comments to all exported functions and types
- Provide usage examples in documentation and tests
- Ensure all cryptographic operations are covered by positive and negative tests, fuzz tests, and benchmarks

## Security
- Never log or expose private keys or secrets in production
- Follow best practices for cryptographic safety and error handling

---

**Use this file to guide all cryptographic code generation for the gopq project.**
