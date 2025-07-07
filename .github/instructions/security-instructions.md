# Security Instructions

These instructions define security best practices for all Go projects, with emphasis on cryptographic and sensitive operations.

## Requirements
- Never log or expose private keys, secrets, or sensitive data in production or test logs.
- Use only NIST- and FIPS 140-3-approved algorithms for cryptography.
- Validate all inputs and handle errors securely.
- Use secure random number generation for all cryptographic operations.
- Review all code for side-channel and timing attack risks.

## Best Practices
- Follow the principle of least privilege for all operations.
- Document all security assumptions and limitations.
- Reference official NIST and FIPS 140-3 documentation in comments and docs.

---

**Use this file to guide all security-related code and review.**
