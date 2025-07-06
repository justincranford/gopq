
# Post-Quantum (PQ) Safe Cryptography Go Library

**Educational Project: Post-Quantum Safe Utility Cryptography created with Agentic AI**

gopq is an educational project to demonstrate:
1. Post-quantum (PQ) Safe Cryptography; ML-DSA for signing, and ML-PEM for encryption.
2. Agentic AI for rapid development; vscode + Copilot instructions and prompts.

## Overview

gopq provides reusable Go functions for NIST PQC algorithms, Kyber KEM and ML-DSA.

The implementation uses the [Cloudflare CIRCL](https://github.com/cloudflare/circl) library.

> **Note:** gopq is for demonstration and educational purposes only. Do not use in production.

## Installation

```
go get github.com/cloudflare/circl@latest
```

Clone or vendor this repository as needed:

```
git clone https://github.com/your-org/gopq.git
```

## Usage

### Key Generation

```go
import "github.com/cloudflare/circl/kem/kyber/kyber1024"
import "your-module-path/internal/crypto/pq"

// Generate a random Kyber1024 KEM keypair
key, err := pq.GenerateMLKEMKeyPair()
if err != nil {
    // handle error
}
```

### Key Serialization

```go
pubBytes, _ := pq.MarshalPublicKey(key.PublicKey)
privBytes, _ := pq.MarshalPrivateKey(key.PrivateKey)

// Deserialize
pub, _ := pq.UnmarshalPublicKey(pubBytes)
priv, _ := pq.UnmarshalPrivateKey(privBytes)
```

### Encapsulation/Decapsulation

```go
// Encapsulate a shared secret
sharedSecret, ciphertext, err := pq.MLKEMEncapsulate(key.PublicKey)

// Decapsulate the shared secret
recoveredSecret, err := pq.MLKEMDecapsulate(key.PrivateKey, ciphertext)
```

### Deterministic KATs (for test vectors)

```go
// Deterministic keypair (for KATs)
seed := make([]byte, kyber1024.Scheme().SeedSize())
detKey, err := pq.GenerateDeterministicMLKEMKeyPair(seed)

// Deterministic encapsulation (for KATs)
encSeed := make([]byte, kyber1024.Scheme().EncapsulationSeedSize())
shared, ct, err := pq.MLKEMEncapsulateDeterministic(detKey.PublicKey, encSeed)
```

## Testing

Run all tests and benchmarks:

```
go test -v -bench=. ./internal/crypto/pq
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

## R&D Persona Prompts

For role-specific guidance and collaboration, see the following persona prompts in `.github/prompts/`:

- [Product Manager](.github/prompts/prompt-pm.md)
- [Senior Staff Developer](.github/prompts/prompt-dev.md)
- [QA Automation Developer](.github/prompts/prompt-qa.md)
- [Documentation Developer](.github/prompts/prompt-doc.md)
- [Release Manager](.github/prompts/prompt-release.md)
- [Marketing VP](.github/prompts/prompt-marketing.md)
- [Application Security Analyst](.github/prompts/prompt-sec.md)
