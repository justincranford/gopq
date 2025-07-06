
# Cryptography Domain Instructions

These instructions define cryptography-specific standards and best practices for Go projects, including but not limited to gopq.

## Requirements
- Use only NIST/FIPS-approved post-quantum algorithms (ML-KEM, ML-DSA, Kyber, Dilithium, etc.).
- Never log or expose private keys or secrets in production or test logs.
- Use secure random number generation for all cryptographic operations.
- Reference official NIST/FIPS documentation in comments and documentation.
- All cryptographic code must be covered by positive and negative tests, fuzz tests, and benchmarks.
- All code must pass `golangci-lint run --fix` and `gofumpt -l -w .` before commit.

## Best Practices
- Use descriptive variable names and wrap all errors for traceability.
- Document all security assumptions and limitations.
- Follow best practices for cryptographic safety and error handling.

---

**Use this file to guide all cryptographic code generation for the gopq project.**
