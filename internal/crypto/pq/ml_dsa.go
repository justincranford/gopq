// Package pq provides post-quantum cryptographic utilities for ML-DSA and ML-KEM.
package pq

import (
	"crypto"
	"crypto/rand"
	"fmt"
	"runtime/debug"

	"github.com/cloudflare/circl/sign/mldsa/mldsa87"
)

// MLDSAKeyPair represents a key pair for ML-DSA.
type MLDSAKeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}

// GenerateMLDSAKeyPair generates a new ML-DSA key pair using CIRCL ML-DSA-87.
func GenerateMLDSAKeyPair() (keyPair *MLDSAKeyPair, keyGenerationError error) {
	defer func() {
		if recoveredPanic := recover(); recoveredPanic != nil {
			keyGenerationError = fmt.Errorf("panic in GenerateMLDSAKeyPair: %v\n%s", recoveredPanic, debug.Stack())
		}
	}()
	publicKey, privateKey, keyGenerationError := mldsa87.GenerateKey(rand.Reader)
	if keyGenerationError != nil {
		return nil, fmt.Errorf("mldsa87.GenerateKey: %w", keyGenerationError)
	}
	publicKeyBytes, publicKeyMarshalError := publicKey.MarshalBinary()
	if publicKeyMarshalError != nil {
		return nil, fmt.Errorf("publicKey.MarshalBinary: %w", publicKeyMarshalError)
	}
	privateKeyBytes, privateKeyMarshalError := privateKey.MarshalBinary()
	if privateKeyMarshalError != nil {
		return nil, fmt.Errorf("privateKey.MarshalBinary: %w", privateKeyMarshalError)
	}
	fmt.Printf("GenerateMLDSAKeyPair: publicKeyBytes len=%d, privateKeyBytes len=%d\n", len(publicKeyBytes), len(privateKeyBytes))
	keyPair = &MLDSAKeyPair{
		PublicKey:  publicKeyBytes,
		PrivateKey: privateKeyBytes,
	}
	return keyPair, nil
}

// DeriveMLDSAKeyPair deterministically derives a new ML-DSA key pair using CIRCL ML-DSA-87 with seed size 32-bytes.
func DeriveMLDSAKeyPair(seed *[mldsa87.SeedSize]byte) (keyPair *MLDSAKeyPair, keyGenerationError error) {
	defer func() {
		if recoveredPanic := recover(); recoveredPanic != nil {
			keyGenerationError = fmt.Errorf("panic in DeriveMLDSAKeyPair: %v\n%s", recoveredPanic, debug.Stack())
		}
	}()
	publicKey, privateKey := mldsa87.NewKeyFromSeed(seed)
	publicKeyBytes, publicKeyMarshalError := publicKey.MarshalBinary()
	if publicKeyMarshalError != nil {
		return nil, fmt.Errorf("publicKey.MarshalBinary: %w", publicKeyMarshalError)
	}
	privateKeyBytes, privateKeyMarshalError := privateKey.MarshalBinary()
	if privateKeyMarshalError != nil {
		return nil, fmt.Errorf("privateKey.MarshalBinary: %w", privateKeyMarshalError)
	}
	fmt.Printf("DeriveMLDSAKeyPair: publicKeyBytes len=%d, privateKeyBytes len=%d\n", len(publicKeyBytes), len(privateKeyBytes))
	keyPair = &MLDSAKeyPair{
		PublicKey:  publicKeyBytes,
		PrivateKey: privateKeyBytes,
	}
	return keyPair, nil
}

// MLDSASign signs a message using the ML-DSA private key (CIRCL ML-DSA-87).
func MLDSASign(privateKeyBytes []byte, messageBytes []byte) (signatureBytes []byte, signError error) {
	defer func() {
		if recoveredPanic := recover(); recoveredPanic != nil {
			signError = fmt.Errorf("panic in MLDSASign: %v\n%s", recoveredPanic, debug.Stack())
		}
	}()
	var privateKey mldsa87.PrivateKey
	if unmarshalError := privateKey.UnmarshalBinary(privateKeyBytes); unmarshalError != nil {
		fmt.Printf("MLDSASign: privateKey.UnmarshalBinary failed: %v\n", unmarshalError)
		fmt.Printf("privateKeyBytes len: %d\n", len(privateKeyBytes))
		return nil, fmt.Errorf("privateKey.UnmarshalBinary failed: %w", unmarshalError)
	}
	// Use crypto.Hash(0) as required by the CIRCL ML-DSA-87 implementation for opts
	signatureBytes, signError = privateKey.Sign(nil, messageBytes, crypto.Hash(0))
	if signError != nil {
		fmt.Printf("MLDSASign: privateKey.Sign failed: %v\n", signError)
		fmt.Printf("messageBytes len: %d\n", len(messageBytes))
		return nil, fmt.Errorf("privateKey.Sign failed: %w", signError)
	}
	fmt.Printf("MLDSASign: messageBytes len=%d, signatureBytes len=%d\n", len(messageBytes), len(signatureBytes))
	return signatureBytes, nil
}

// MLDSAVerify verifies an ML-DSA signature using CIRCL ML-DSA-87.
func MLDSAVerify(publicKeyBytes []byte, messageBytes []byte, signatureBytes []byte) (isSignatureValid bool, verifyError error) {
	defer func() {
		if recoveredPanic := recover(); recoveredPanic != nil {
			verifyError = fmt.Errorf("panic in MLDSAVerify: %v\n%s", recoveredPanic, debug.Stack())
		}
	}()
	var publicKey mldsa87.PublicKey
	if unmarshalError := publicKey.UnmarshalBinary(publicKeyBytes); unmarshalError != nil {
		fmt.Printf("MLDSAVerify: publicKey.UnmarshalBinary failed: %v\n", unmarshalError)
		fmt.Printf("publicKeyBytes len: %d\n", len(publicKeyBytes))
		return false, fmt.Errorf("publicKey.UnmarshalBinary failed: %w", unmarshalError)
	}
	isSignatureValid = mldsa87.Verify(&publicKey, messageBytes, nil, signatureBytes)
	fmt.Printf("MLDSAVerify: messageBytes len=%d, signatureBytes len=%d, isSignatureValid=%v\n", len(messageBytes), len(signatureBytes), isSignatureValid)
	return isSignatureValid, nil
}
