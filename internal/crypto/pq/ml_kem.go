// Package pq provides post-quantum cryptographic utilities for ML-DSA and ML-KEM.
package pq

import (
	"errors"
	"runtime/debug"

	"github.com/cloudflare/circl/kem"
	"github.com/cloudflare/circl/kem/kyber/kyber1024"
)

// MLKEMKeyPair represents a key pair for Kyber1024 KEM.
type MLKEMKeyPair struct {
	PublicKey  kem.PublicKey
	PrivateKey kem.PrivateKey
}

// GenerateDeterministicMLKEMKeyPair generates a Kyber1024 KEM key pair from a seed (for KATs).
// The seed must be of length kyber1024.Scheme().SeedSize().
func GenerateDeterministicMLKEMKeyPair(seed []byte) (*MLKEMKeyPair, error) {
	defer func() {
		if r := recover(); r != nil {
			println("panic in GenerateDeterministicMLKEMKeyPair:", r)
			debug.PrintStack()
		}
	}()
	if len(seed) != kyber1024.Scheme().SeedSize() {
		return nil, errors.New("invalid seed size")
	}
	publicKey, privateKey := kyber1024.Scheme().DeriveKeyPair(seed)
	return &MLKEMKeyPair{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}, nil
}

// GenerateMLKEMKeyPair generates a new Kyber1024 KEM key pair.
func GenerateMLKEMKeyPair() (*MLKEMKeyPair, error) {
	defer func() {
		if r := recover(); r != nil {
			println("panic in GenerateMLKEMKeyPair:", r)
			debug.PrintStack()
		}
	}()
	publicKey, privateKey, err := kyber1024.Scheme().GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	return &MLKEMKeyPair{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}, nil
}

// MarshalPublicKey serializes a Kyber1024 public key to bytes.
func MarshalPublicKey(pk kem.PublicKey) ([]byte, error) {
	defer func() {
		if r := recover(); r != nil {
			println("panic in MarshalPublicKey:", r)
			debug.PrintStack()
		}
	}()
	return pk.MarshalBinary()
}

// UnmarshalPublicKey deserializes bytes into a Kyber1024 public key.
func UnmarshalPublicKey(data []byte) (kem.PublicKey, error) {
	defer func() {
		if r := recover(); r != nil {
			println("panic in UnmarshalPublicKey:", r)
			debug.PrintStack()
		}
	}()
	return kyber1024.Scheme().UnmarshalBinaryPublicKey(data)
}

// MarshalPrivateKey serializes a Kyber1024 private key to bytes.
func MarshalPrivateKey(sk kem.PrivateKey) ([]byte, error) {
	defer func() {
		if r := recover(); r != nil {
			println("panic in MarshalPrivateKey:", r)
			debug.PrintStack()
		}
	}()
	return sk.MarshalBinary()
}

// UnmarshalPrivateKey deserializes bytes into a Kyber1024 private key.
func UnmarshalPrivateKey(data []byte) (kem.PrivateKey, error) {
	defer func() {
		if r := recover(); r != nil {
			println("panic in UnmarshalPrivateKey:", r)
			debug.PrintStack()
		}
	}()
	return kyber1024.Scheme().UnmarshalBinaryPrivateKey(data)
}

// MLKEMEncapsulate encapsulates a shared secret using the Kyber1024 public key.
func MLKEMEncapsulate(publicKey kem.PublicKey) (ciphertext []byte, sharedSecret []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			println("panic in MLKEMEncapsulate:", r)
			debug.PrintStack()
		}
	}()
	if publicKey == nil {
		return nil, nil, errors.New("invalid public key")
	}
	ciphertext, sharedSecret, err = kyber1024.Scheme().Encapsulate(publicKey)
	return ciphertext, sharedSecret, err
}

// MLKEMDecapsulate decapsulates a shared secret using the Kyber1024 private key.
func MLKEMDecapsulate(privateKey kem.PrivateKey, ciphertext []byte) ([]byte, error) {
	defer func() {
		if r := recover(); r != nil {
			println("panic in MLKEMDecapsulate:", r)
			debug.PrintStack()
		}
	}()
	if privateKey == nil || len(ciphertext) == 0 {
		return nil, errors.New("invalid input")
	}
	sharedSecret, err := kyber1024.Scheme().Decapsulate(privateKey, ciphertext)
	if err != nil || sharedSecret == nil || len(sharedSecret) == 0 {
		return nil, errors.New("decapsulation failed")
	}
	// Optionally, check for all-zero shared secret (could indicate failure)
	allZero := true
	for _, b := range sharedSecret {
		if b != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return nil, errors.New("decapsulation failed: all-zero shared secret")
	}
	return sharedSecret, nil
}

// MLKEMEncapsulateDeterministic encapsulates a shared secret using the Kyber1024 public key and a seed (for KATs).
// The seed must be of length kyber1024.Scheme().EncapsulationSeedSize().
func MLKEMEncapsulateDeterministic(publicKey kem.PublicKey, seed []byte) (ciphertext []byte, sharedSecret []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			println("panic in MLKEMEncapsulateDeterministic:", r)
			debug.PrintStack()
		}
	}()
	if publicKey == nil {
		return nil, nil, errors.New("invalid public key")
	}
	if len(seed) != kyber1024.Scheme().EncapsulationSeedSize() {
		return nil, nil, errors.New("invalid encapsulation seed size")
	}
	ciphertext, sharedSecret, err = kyber1024.Scheme().EncapsulateDeterministically(publicKey, seed)
	return ciphertext, sharedSecret, err
}
