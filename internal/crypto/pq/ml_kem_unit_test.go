package pq

import (
	"crypto/subtle"
	"testing"
)

func TestGenerateMLKEMKeyPair(t *testing.T) {
	keyPair, err := GenerateMLKEMKeyPair()
	if err != nil {
		t.Fatalf("failed to generate ML-KEM key pair: %v", err)
	}
	if keyPair.PublicKey == nil || keyPair.PrivateKey == nil {
		t.Error("key pair should not be nil")
	}
	publicKeyBytes, err := MarshalPublicKey(keyPair.PublicKey)
	if err != nil {
		t.Fatalf("marshal public key failed: %v", err)
	}
	privateKeyBytes, err := MarshalPrivateKey(keyPair.PrivateKey)
	if err != nil {
		t.Fatalf("marshal private key failed: %v", err)
	}
	publicKey, err := UnmarshalPublicKey(publicKeyBytes)
	if err != nil {
		t.Fatalf("unmarshal public key failed: %v", err)
	}
	privateKey, err := UnmarshalPrivateKey(privateKeyBytes)
	if err != nil {
		t.Fatalf("unmarshal private key failed: %v", err)
	}
	if publicKey == nil || privateKey == nil {
		t.Error("unmarshaled keys should not be nil")
	}
}

func TestMLKEMEncapsulateAndDecapsulate(t *testing.T) {
	keyPair, _ := GenerateMLKEMKeyPair()
	ciphertext, sharedSecret, err := MLKEMEncapsulate(keyPair.PublicKey)
	if err != nil {
		t.Fatalf("encapsulation failed: %v", err)
	}
	sharedSecret2, err := MLKEMDecapsulate(keyPair.PrivateKey, ciphertext)
	if err != nil {
		t.Fatalf("decapsulation failed: %v", err)
	}
	if subtle.ConstantTimeCompare(sharedSecret, sharedSecret2) == 0 {
		t.Error("shared secrets were expected to match, decapculate with private key and ciphertext should have returned a same shared secret as encapculate with public key")
	}
}

func TestMLKEMEncapsulateWithInvalidKey(t *testing.T) {
	_, _, err := MLKEMEncapsulate(nil)
	if err == nil {
		t.Error("expected error for nil public key")
	}
}

func TestMLKEMDecapsulateWithInvalidKey(t *testing.T) {
	ciphertext, _, _ := MLKEMEncapsulate(nil)
	_, err := MLKEMDecapsulate(nil, ciphertext)
	if err == nil {
		t.Error("expected error for nil private key")
	}
}

func TestMLKEMDecapsulateWithTamperedCiphertext(t *testing.T) {
	keyPair, _ := GenerateMLKEMKeyPair()
	ciphertext, _, _ := MLKEMEncapsulate(keyPair.PublicKey)
	sharedSecret, err := MLKEMDecapsulate(keyPair.PrivateKey, ciphertext)
	if err != nil {
		t.Fatalf("decapsulation failed for original ciphertext: %v", err)
	}
	if len(ciphertext) > 0 {
		ciphertext[0] ^= 0xFF // tamper ciphertext
	}
	sharedSecretTampered, err := MLKEMDecapsulate(keyPair.PrivateKey, ciphertext)
	if err != nil {
		t.Logf("decapsulation failed for tampered ciphertext (acceptable): %v", err)
		return
	}
	if subtle.ConstantTimeCompare(sharedSecret, sharedSecretTampered) == 1 {
		t.Error("shared secrets were expected to be different, decapsulated with tampered ciphertext should have returned a different shared secret")
	}
}

func TestMLKEMDecapsulateWithWrongKey(t *testing.T) {
	keyPair1, _ := GenerateMLKEMKeyPair()
	keyPair2, _ := GenerateMLKEMKeyPair()
	ciphertext, sharedSecret, _ := MLKEMEncapsulate(keyPair1.PublicKey)
	sharedWrong, err := MLKEMDecapsulate(keyPair2.PrivateKey, ciphertext)
	if err != nil {
		t.Logf("decapsulation failed for wrong private key (acceptable): %v", err)
		return
	}
	if subtle.ConstantTimeCompare(sharedSecret, sharedWrong) == 1 {
		t.Error("shared secrets were expected to be different, decapsulated with wrong private key should have returned a different shared secret")
	}
}
