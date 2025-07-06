
# Post-Quantum (PQ) Safe Cryptography Go Library

**Educational Project: Post-Quantum Safe Utility Cryptography created with Agentic AI**

gopq is an educational project to demonstrate:
1. Post-quantum (PQ) Safe Cryptography; ML-DSA for signing, and ML-PEM for encryption.
2. Agentic AI for rapid development; vscode + Copilot instructions and prompts.

## Overview

gopq provides reusable Go functions for NIST PQC algorithms, Kyber KEM and ML-DSA.

> **Note:** gopq is for demonstration and educational purposes only. Do not use in production.

The implementation uses the [Cloudflare CIRCL](https://github.com/cloudflare/circl) library.

## Installation

```
go get github.com/cloudflare/circl@latest
```

Clone or vendor this repository as needed:

```
git clone https://github.com/your-org/gopq.git
```


## Usage

### ML-KEM (Kyber KEM) Example

```go
import "github.com/cloudflare/circl/kem/kyber/kyber1024"
import "gopq/pq"

// Generate a random Kyber1024 KEM keypair
mlkemKey, err := pq.GenerateMLKEMKeyPair()
if err != nil {
    // handle error
}

// Serialize and deserialize keys
pubBytes, _ := pq.MarshalPublicKey(mlkemKey.PublicKey)
privBytes, _ := pq.MarshalPrivateKey(mlkemKey.PrivateKey)
pub, _ := pq.UnmarshalPublicKey(pubBytes)
priv, _ := pq.UnmarshalPrivateKey(privBytes)

// Encapsulate a shared secret
ciphertext, sharedSecret, err := pq.MLKEMEncapsulate(pub)

// Decapsulate the shared secret
recoveredSecret, err := pq.MLKEMDecapsulate(priv, ciphertext)

// Deterministic keypair (for KATs)
seed := make([]byte, kyber1024.Scheme().SeedSize())
detKey, err := pq.GenerateDeterministicMLKEMKeyPair(seed)

// Deterministic encapsulation (for KATs)
encSeed := make([]byte, kyber1024.Scheme().EncapsulationSeedSize())
ct, shared, err := pq.MLKEMEncapsulateDeterministic(detKey.PublicKey, encSeed)
```

### ML-DSA (ML-DSA-87) Example

```go
import "github.com/cloudflare/circl/sign/mldsa/mldsa87"
import "gopq/pq"

// Generate a random ML-DSA keypair
mldsaKey, err := pq.GenerateMLDSAKeyPair()
if err != nil {
    // handle error
}

// Deterministic keypair (from seed)
var seed [mldsa87.SeedSize]byte
detDSAKey, err := pq.DeriveMLDSAKeyPair(&seed)

// Sign a message
message := []byte("hello world")
signature, err := pq.MLDSASign(mldsaKey.PrivateKey, message)
if err != nil {
    // handle error
}

// Verify a signature
valid, err := pq.MLDSAVerify(mldsaKey.PublicKey, message, signature)
if err != nil {
    // handle error
}
if !valid {
    // signature invalid
}
```

## Testing

Run all tests and benchmarks:

```
go test -v -bench=. ./pq
```

## Security Notes

- This library is for demonstration and educational use only.
- For production, use vetted libraries and follow NIST/FIPS guidance.
- Never log or expose private keys or shared secrets in production.
- **Limitations:**
  - The Kyber KEM decapsulation (via CIRCL) may not always return an error for tampered ciphertext or wrong private key, depending on the underlying library's behavior. Always validate shared secrets and handle errors securely in production.


## References

- [Cloudflare CIRCL](https://github.com/cloudflare/circl)
- [NIST PQC Standardization](https://csrc.nist.gov/projects/post-quantum-cryptography)
- [Kyber Specification](https://pq-crystals.org/kyber/)

## Agentic AI Usage in This Project

This project was developed using an Agentic AI workflow, leveraging VS Code, GitHub Copilot, and custom instruction and prompt files to enforce standards, automate code generation, and maintain project quality.

### Key Agentic AI Artifacts

- [Copilot/AI Instructions](.github/instructions/copilot-instructions.md): Defines how Copilot and other AI tools must operate in this project.
- [Commit Workflow Instructions](.github/instructions/commit-instructions.md): Enforces commit and workflow rules, including PowerShell command usage.
- [Coding Standards](.github/instructions/coding-instructions.md): Project-wide code quality and style requirements.
- [Go Language Standards](.github/instructions/go-instructions.md): All Go-specific requirements and best practices.
- [Security Standards](.github/instructions/security-instructions.md): Security and cryptography requirements.
- [Documentation Standards](.github/instructions/doc-instructions.md): Documentation and GoDoc requirements.
- [Test Instructions](.github/instructions/test-instructions.md): Unit, fuzz, and benchmark test requirements.
- [Project Instructions](.github/instructions/project-instructions.md): Project-specific rules and meta-guidance.
- [Persona Prompts](.github/prompts/): Role-specific prompts for developers, QA, security, product, and more.

These files ensure that all contributors and AI tools follow the same standards, enabling rapid, consistent, and high-quality development.

## R&D Persona Prompts

For role-specific guidance and collaboration, see the following persona prompts in `.github/prompts/`:

- [Product Manager](.github/prompts/prompt-pm.md)
- [Senior Staff Developer](.github/prompts/prompt-dev.md)
- [QA Automation Developer](.github/prompts/prompt-qa.md)
- [Documentation Developer](.github/prompts/prompt-doc.md)
- [Release Manager](.github/prompts/prompt-release.md)
- [Marketing VP](.github/prompts/prompt-marketing.md)
- [Application Security Analyst](.github/prompts/prompt-sec.md)
