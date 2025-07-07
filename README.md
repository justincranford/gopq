
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
git clone https://github.com/justincranford/gopq.git
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

### Generating and Viewing GoDoc Documentation

To generate and view GoDoc documentation locally:

1. Install the godoc tool (if not already installed):
   ```
   go install golang.org/x/tools/cmd/godoc@latest
   ```
2. Start the documentation server:
   ```
   godoc -http=:6060
   ```
3. Open your browser and go to [http://localhost:6060/pkg/](http://localhost:6060/pkg/) to browse the documentation for this project and its packages.

### Security Notes

- This library is for demonstration and educational use only.
- For production, use vetted libraries and follow NIST and FIPS 140-3 guidance.
- Never log or expose private keys or shared secrets in production.

### References

- [Cloudflare CIRCL](https://github.com/cloudflare/circl)
- [NIST PQC Standardization](https://csrc.nist.gov/projects/post-quantum-cryptography)
- [Kyber Specification](https://pq-crystals.org/kyber/)

## Agentic AI Usage in This Project

This project was developed using an Agentic AI workflow, leveraging VS Code, GitHub Copilot, and custom instruction and prompt files. The main configuration file is `settings.json`. It lists the instruction and prompt files to be used.

- [.vscode/settings.json](.vscode/settings.json)

### General R&D Instruction Files


Instruction files ensure that all human contributors and AI personas follow the same standards. They are in the [.github/instructions/](.github/instructions/) directory:

AI

- [copilot-instructions.md](.github/instructions/copilot-instructions.md)

Development

- [commit-instructions.md](.github/instructions/commit-instructions.md)
- [coding-instructions.md](.github/instructions/coding-instructions.md)
- [go-instructions.md](.github/instructions/go-instructions.md)
- [security-instructions.md](.github/instructions/security-instructions.md)
- [doc-instructions.md](.github/instructions/doc-instructions.md)
- [project-instructions.md](.github/instructions/project-instructions.md)

Testing

- [test-instructions.md](.github/instructions/test-instructions.md)
- [test-fuzz-instructions.md](.github/instructions/test-fuzz-instructions.md)
- [test-bench-instructions.md](.github/instructions/test-bench-instructions.md)

Design intent is for all instruction files to be generic and reusable for other R&D projects, except for project-instructions, which is project-specific.

### Persona Prompt Files


All of these files are persona prompts in the [.github/prompts/](.github/prompts/) directory:

- [prompt-dev.md](.github/prompts/prompt-dev.md)
- [prompt-qa.md](.github/prompts/prompt-qa.md)
- [prompt-pm.md](.github/prompts/prompt-pm.md)
- [prompt-sec.md](.github/prompts/prompt-sec.md)
- [prompt-doc.md](.github/prompts/prompt-doc.md)
- [prompt-marketing.md](.github/prompts/prompt-marketing.md)

### Additional Details

Instruction files define standards and rules that apply to both human contributors and AI agents.

When a human contributor interacts with Agentic AI:
- Using Copilot Chat, the AI agent follows all instructions in a general context (no specific persona).
- Using a Persona Prompt, the AI agent follows all instructions, but adopts the perspective and goals of the selected persona.

AI is most effective when given specific context.

Using Copilot Chat is good, but sometimes you get mixed results. When that happens, switch to a persona (e.g., Dev, QA, PM) to get more focused results and expertise.
