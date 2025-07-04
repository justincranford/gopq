// Package pq provides post-quantum cryptographic utilities for ML-DSA and ML-KEM.
package pq

import (
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
func GenerateMLDSAKeyPair() (*MLDSAKeyPair, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic in GenerateMLDSAKeyPair: %v\n%s", r, debug.Stack())
		}
	}()
	pk, sk, err := mldsa87.GenerateKey(nil)
	if err != nil {
		return nil, fmt.Errorf("mldsa87.GenerateKey: %w", err)
	}
	pub, err := pk.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("pk.MarshalBinary: %w", err)
	}
	priv, err := sk.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("sk.MarshalBinary: %w", err)
	}
	return &MLDSAKeyPair{
		PublicKey:  pub,
		PrivateKey: priv,
	}, nil
}

// MLDSASign signs a message using the ML-DSA private key (CIRCL ML-DSA-87).
func MLDSASign(privateKey []byte, message []byte) (sig []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic in MLDSASign: %v\n%s", r, debug.Stack())
			sig = nil
		}
	}()
	var sk mldsa87.PrivateKey
	if err := sk.UnmarshalBinary(privateKey); err != nil {
		fmt.Printf("MLDSASign: sk.UnmarshalBinary failed: %v\n", err)
		fmt.Printf("privateKey len: %d\n", len(privateKey))
		return nil, err
	}
	sig, err = sk.Sign(nil, message, nil)
	if err != nil {
		fmt.Printf("MLDSASign: sk.Sign failed: %v\n", err)
		fmt.Printf("message len: %d\n", len(message))
		return nil, err
	}
	fmt.Printf("MLDSASign: sig len: %d\n", len(sig))
	return sig, nil
}

// MLDSAVerify verifies an ML-DSA signature using CIRCL ML-DSA-87.
func MLDSAVerify(publicKey []byte, message []byte, signature []byte) (verified bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic in MLDSAVerify: %v\n%s", r, debug.Stack())
			verified = false
		}
	}()
	var pk mldsa87.PublicKey
	if err := pk.UnmarshalBinary(publicKey); err != nil {
		fmt.Printf("MLDSAVerify: pk.UnmarshalBinary failed: %v\n", err)
		fmt.Printf("publicKey len: %d\n", len(publicKey))
		return false, fmt.Errorf("pk.UnmarshalBinary failed: %w", err)
	}
	valid := mldsa87.Verify(&pk, message, nil, signature)
	fmt.Printf("MLDSAVerify: valid=%v, message len=%d, signature len=%d\n", valid, len(message), len(signature))
	return valid, nil
}
