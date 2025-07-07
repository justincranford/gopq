
# Post-Quantum (PQ) Safe Cryptography Go Library

**Educational Project: Post-Quantum Safe Utility Cryptography created with Agentic AI**


gopq is an educational project to demonstrate:
- [Post-quantum Cryptography](#post-quantum-cryptography)
- [Agentic AI Usage in This Project](#agentic-ai-usage-in-this-project)

## Post-quantum Cryptography

gopq provides reusable Go functions for PQC algorithms; ML-DSA for signing, and ML-KEM for encryption.

> **Note:** gopq is for demonstration and educational purposes only. Do not use in production.

The implementation uses the [Cloudflare CIRCL](https://github.com/cloudflare/circl) library.

### Installation

```
go get github.com/cloudflare/circl@latest
```

Clone or vendor this repository as needed:

```
git clone https://github.com/your-org/gopq.git
```


### Usage

#### ML-DSA (ML-DSA-87) Example

```
import "gopq/pq"

// Generate a random ML-DSA keypair
mldsaKey, err := pq.GenerateMLDSAKeyPair()
if err != nil {
    // handle error
}

// Deterministic keypair (from seed)
var seed [48]byte // 48 is the seed size for ML-DSA-87
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

#### ML-KEM (Kyber KEM) Example

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

### Testing

Run all tests and benchmarks:

```
go test -v -bench=. ./pq
```

### Security Notes

- This library is for demonstration and educational use only.
- For production, use vetted libraries and follow NIST and FIPS 140-3 guidance.
- Never log or expose private keys or shared secrets in production.

### References

- [Cloudflare CIRCL](https://github.com/cloudflare/circl)
- [NIST PQC Standardization](https://csrc.nist.gov/projects/post-quantum-cryptography)
- [Kyber Specification](https://pq-crystals.org/kyber/)

## Agentic AI Usage in This Project


This project was developed using an Agentic AI workflow, leveraging VS Code, GitHub Copilot, and custom instruction and prompt files to enforce standards, automate code generation, and maintain project quality.

#### Relationship of settings.json, instructions, and prompts

The `.vscode/settings.json` file controls how Copilot and other AI tools use the instruction and prompt files. All instruction and prompt files must be referenced in settings.json to ensure AI tools follow the correct standards and workflows. Prompts are always generic and must defer to the instruction files for all requirements and standards. When updating or adding instruction or prompt files, always update settings.json and validate the references.

**Key files and directories:**

- [settings.json](.vscode/settings.json)
- [.github/instructions/](.github/instructions/)
- [.github/prompts/](.github/prompts/)


### Key Agentic AI Artifacts

- [copilot-instructions.md](.github/instructions/copilot-instructions.md)
- [commit-instructions.md](.github/instructions/commit-instructions.md)
- [coding-instructions.md](.github/instructions/coding-instructions.md)
- [go-instructions.md](.github/instructions/go-instructions.md)
- [security-instructions.md](.github/instructions/security-instructions.md)
- [doc-instructions.md](.github/instructions/doc-instructions.md)
- [test-instructions.md](.github/instructions/test-instructions.md)
- [project-instructions.md](.github/instructions/project-instructions.md)
- [prompts](.github/prompts/)


These files ensure that all contributors and AI tools follow the same standards, enabling rapid, consistent, and high-quality development.


### R&D Persona Prompts

For role-specific guidance and collaboration, see the following persona prompt files:

- [prompt-pm.md](.github/prompts/prompt-pm.md)
- [prompt-dev.md](.github/prompts/prompt-dev.md)
- [prompt-qa.md](.github/prompts/prompt-qa.md)
- [prompt-doc.md](.github/prompts/prompt-doc.md)
- [prompt-release.md](.github/prompts/prompt-release.md)
- [prompt-marketing.md](.github/prompts/prompt-marketing.md)
- [prompt-sec.md](.github/prompts/prompt-sec.md)
