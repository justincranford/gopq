// Package pq provides post-quantum cryptographic utilities for ML-DSA and ML-KEM.
package pq

import (
	"errors"

	"github.com/cloudflare/circl/kem"
	"github.com/cloudflare/circl/kem/kyber/kyber1024"
)

// MLKEMKeyPair represents a key pair for Kyber1024 KEM.
type MLKEMKeyPair struct {
	PublicKey  kem.PublicKey
	PrivateKey kem.PrivateKey
}

// GenerateMLKEMKeyPair generates a new Kyber1024 KEM key pair.
func GenerateMLKEMKeyPair() (*MLKEMKeyPair, error) {
	pk, sk, err := kyber1024.Scheme().GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	return &MLKEMKeyPair{PublicKey: pk, PrivateKey: sk}, nil
}

// MarshalPublicKey serializes a Kyber1024 public key to bytes.
func MarshalPublicKey(pk kem.PublicKey) ([]byte, error) {
	return pk.MarshalBinary()
}

// UnmarshalPublicKey deserializes bytes into a Kyber1024 public key.
func UnmarshalPublicKey(data []byte) (kem.PublicKey, error) {
	return kyber1024.Scheme().UnmarshalBinaryPublicKey(data)
}

// MarshalPrivateKey serializes a Kyber1024 private key to bytes.
func MarshalPrivateKey(sk kem.PrivateKey) ([]byte, error) {
	return sk.MarshalBinary()
}

// UnmarshalPrivateKey deserializes bytes into a Kyber1024 private key.
func UnmarshalPrivateKey(data []byte) (kem.PrivateKey, error) {
	return kyber1024.Scheme().UnmarshalBinaryPrivateKey(data)
}

// MLKEMEncapsulate encapsulates a shared secret using the Kyber1024 public key.
func MLKEMEncapsulate(publicKey kem.PublicKey) (sharedSecret []byte, ciphertext []byte, err error) {
	if publicKey == nil {
		return nil, nil, errors.New("invalid public key")
	}
	ct, ss, err := kyber1024.Scheme().Encapsulate(publicKey)
	return ss, ct, err
}

// MLKEMDecapsulate decapsulates a shared secret using the Kyber1024 private key.
func MLKEMDecapsulate(privateKey kem.PrivateKey, ciphertext []byte) ([]byte, error) {
	if privateKey == nil || len(ciphertext) == 0 {
		return nil, errors.New("invalid input")
	}
	ss, err := kyber1024.Scheme().Decapsulate(privateKey, ciphertext)
	return ss, err
}
