package pq

import (
	"testing"
)

func TestGenerateMLDSAKeyPair(t *testing.T) {
	key, err := GenerateMLDSAKeyPair()
	if err != nil {
		t.Fatalf("failed to generate ML-DSA key pair: %v", err)
	}
	if len(key.PublicKey) == 0 || len(key.PrivateKey) == 0 {
		t.Error("key pair should not be empty")
	}
}

func TestMLDSASignAndVerify(t *testing.T) {
	key, err := GenerateMLDSAKeyPair()
	if err != nil {
		t.Fatalf("failed to generate ML-DSA key pair: %v", err)
	}
	msg := []byte("test message")
	sig, err := MLDSASign(key.PrivateKey, msg)
	if err != nil {
		t.Fatalf("signing failed: %v", err)
	}
	isVerify, err := MLDSAVerify(key.PublicKey, msg, sig)
	if err != nil {
		t.Fatalf("verifying failed: %v", err)
	}
	if !isVerify {
		t.Error("signature should verify")
	}
}

func FuzzMLDSASignAndVerify(f *testing.F) {
	key, _ := GenerateMLDSAKeyPair()
	f.Fuzz(func(t *testing.T, msg []byte) {
		sig, err := MLDSASign(key.PrivateKey, msg)
		if err != nil {
			t.Skip()
		}
		isVerify, err := MLDSAVerify(key.PublicKey, msg, sig)
		if err != nil {
			t.Fatalf("verifying failed: %v", err)
		}
		if !isVerify {
			t.Error("signature should verify for fuzzed input")
		}
	})
}

func BenchmarkMLDSASign(b *testing.B) {
	key, err := GenerateMLDSAKeyPair()
	if err != nil {
		b.Fatalf("failed to generate ML-DSA key pair: %v", err)
	}
	msg := []byte("benchmark message")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := MLDSASign(key.PrivateKey, msg)
		if err != nil {
			b.Fatalf("signing failed: %v", err)
		}
	}
}

func BenchmarkMLDSAVerify(b *testing.B) {
	key, err := GenerateMLDSAKeyPair()
	if err != nil {
		b.Fatalf("failed to generate ML-DSA key pair: %v", err)
	}
	msg := []byte("benchmark message")
	sig, err := MLDSASign(key.PrivateKey, msg)
	if err != nil {
		b.Fatalf("signing failed: %v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := MLDSAVerify(key.PublicKey, msg, sig)
		if err != nil {
			b.Fatalf("verifying failed: %v", err)
		}
	}
}

func TestMLDSASignWithInvalidKey(t *testing.T) {
	_, err := MLDSASign([]byte{}, []byte("msg"))
	if err == nil {
		t.Error("expected error for empty private key")
	}
}

func TestMLDSAVerifyWithInvalidKey(t *testing.T) {
	key, err := GenerateMLDSAKeyPair()
	if err != nil {
		t.Fatalf("failed to generate ML-DSA key pair: %v", err)
	}
	msg := []byte("msg")
	sig, err := MLDSASign(key.PrivateKey, msg)
	if err != nil {
		t.Fatalf("signing failed: %v", err)
	}
	invalidPub := []byte{}
	isVerify, err := MLDSAVerify(invalidPub, msg, sig)
	if err == nil {
		t.Error("expected error for invalid public key")
	}
	if isVerify {
		t.Error("verify should fail with invalid public key")
	}
}

func TestMLDSAVerifyWithTamperedSignature(t *testing.T) {
	key, err := GenerateMLDSAKeyPair()
	if err != nil {
		t.Fatalf("failed to generate ML-DSA key pair: %v", err)
	}
	msg := []byte("msg")
	sig, err := MLDSASign(key.PrivateKey, msg)
	if err != nil {
		t.Fatalf("signing failed: %v", err)
	}
	isVerify, err := MLDSAVerify(key.PublicKey, msg, sig)
	if err != nil {
		t.Fatalf("verifying failed: %v", err)
	}
	if !isVerify {
		t.Error("verify should succeed for original signature")
	}
	if len(sig) > 0 {
		sig[0] ^= 0xFF // tamper signature
	}
	isVerify, err = MLDSAVerify(key.PublicKey, msg, sig)
	if err == nil {
		t.Error("expected error for tampered signature")
	}
	if isVerify {
		t.Error("verify should fail with tampered signature")
	}
}
