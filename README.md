
# Post-Quantum (PQ) Safe Cryptography Go Library

**Educational Project: Post-Quantum Safe Utility Cryptography created with Agentic AI**


gopq is an educational project to demonstrate:
- [Post-quantum Cryptography](#post-quantum-cryptography)
- [Agentic AI Usage in This Project](#agentic-ai-usage-in-this-project)

## Post-quantum Cryptography

### Overview

<details>
<summary><strong>Overview</strong></summary>

gopq provides reusable Go functions for PQC algorithms; ML-DSA for signing, and ML-KEM for encryption.

> **Note:** gopq is for demonstration and educational purposes only. Do not use in production.

The implementation uses the [Cloudflare CIRCL](https://github.com/cloudflare/circl) library.

</details>

### Installation

<details>
<summary><strong>Installation</strong></summary>

```
go get github.com/cloudflare/circl@latest
```

Clone or vendor this repository as needed:

```
git clone https://github.com/justincranford/gopq.git
```

</details>

### Usage

<details>
<summary><strong>ML-DSA (ML-DSA-87) Example</strong></summary>

```go
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

</details>

<details>
<summary><strong>ML-KEM (Kyber KEM) Example</strong></summary>

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

</details>


<details>
<summary><strong>Testing</strong></summary>

Run all tests and benchmarks:

```
go test -v -bench=. ./pq
```

</details>


<details>
<summary><strong>Documentation</strong></summary>

To generate and view GoDoc documentation locally:

1. Install the godoc tool (if not already installed):
   ```
   go install golang.org/x/tools/cmd/godoc@latest
   ```
2. Start the documentation server:
   ```
   godoc -http=:6060
   ```
3. Open the documentation link [http://localhost:6060/pkg/gopq/pq](http://localhost:6060/pkg/gopq/pq) in your browser.
   ```
   start http://localhost:6060/pkg/gopq/pq/
   ```

</details>


<details>
<summary><strong>Security Notes</strong></summary>

- This library is for demonstration and educational use only.
- For production, use vetted libraries and follow NIST and FIPS 140-3 guidance.
- Never log or expose private keys or shared secrets in production.

</details>


<details>
<summary><strong>References</strong></summary>

- [Cloudflare CIRCL](https://github.com/cloudflare/circl)
- [NIST PQC Standardization](https://csrc.nist.gov/projects/post-quantum-cryptography)
- [Kyber Specification](https://pq-crystals.org/kyber/)

</details>

## Agentic AI Usage in This Project

This project was developed using an Agentic AI workflow, leveraging:
- VS Code
- GitHub Copilot
- Custom instruction and prompt files

The main configuration file is [.vscode/settings.json](.vscode/settings.json). It specifies the instruction and prompt files to use.


<details>
<summary><strong>Instruction Files</strong></summary>

All instruction files are in the [.github/instructions/](.github/instructions/) directory.

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

Instruction files ensure that all human contributors and AI personas follow the same standards.

Design intent for the instruction files is to be generic and reusable for my other R&D projects. The main exception is project-instructions file, which contains project-specific context.

</details>

### Prompt Files

<details>
<summary><strong>Prompt Files</strong></summary>

All prompt files are in the [.github/prompts/](.github/prompts/) directory.

- [prompt-dev.md](.github/prompts/prompt-dev.md)
- [prompt-qa.md](.github/prompts/prompt-qa.md)
- [prompt-pm.md](.github/prompts/prompt-pm.md)
- [prompt-sec.md](.github/prompts/prompt-sec.md)
- [prompt-doc.md](.github/prompts/prompt-doc.md)
- [prompt-marketing.md](.github/prompts/prompt-marketing.md)

Prompt files are personas for AI Agents to adopt more focused context when performing tasks. See Additional Details.

</details>

### Additional Details

<details>
<summary><strong>Additional Details</strong></summary>

Instruction files define standards and rules that apply to both human contributors and AI agents.

When a human contributor interacts with Agentic AI:
- Using **Copilot Chat**, the AI agent follows all instructions in a general context (no specific persona).
- Using a **Persona Prompt**, the AI agent follows all instructions, but adopts the perspective and goals of the selected persona.

AI effectiveness can increase when given more specific and narrow context.

For example, you can get good but mixed results using Copilot Chat. Switching to a persona (e.g., Dev, QA, PM) can narrow the context, and help AI give you the better expertise and results you were expecting.

</details>
