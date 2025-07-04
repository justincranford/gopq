// Package pq provides post-quantum cryptographic utilities for ML-DSA and ML-KEM.
package pq

import (
	"fmt"

	"github.com/cloudflare/circl/sign/mldsa/mldsa87"
)

// MLDSAKeyPair represents a key pair for ML-DSA.
type MLDSAKeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}

// GenerateMLDSAKeyPair generates a new ML-DSA key pair using CIRCL ML-DSA-87.
func GenerateMLDSAKeyPair() (*MLDSAKeyPair, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic in GenerateMLDSAKeyPair: %v\n", r)
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
func MLDSASign(privateKey []byte, message []byte) ([]byte, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic in MLDSASign: %v\n", r)
		}
	}()
	var sk mldsa87.PrivateKey
	if err := sk.UnmarshalBinary(privateKey); err != nil {
		return nil, err
	}
	sig, err := sk.Sign(nil, message, nil)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

// MLDSAVerify verifies an ML-DSA signature using CIRCL ML-DSA-87.
func MLDSAVerify(publicKey []byte, message []byte, signature []byte) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic in MLDSAVerify: %v\n", r)
		}
	}()
	var pk mldsa87.PublicKey
	if err := pk.UnmarshalBinary(publicKey); err != nil {
		return false
	}
	return mldsa87.Verify(&pk, message, nil, signature)
}
