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
	key, _ := GenerateMLDSAKeyPair()
	msg := []byte("test message")
	sig, err := MLDSASign(key.PrivateKey, msg)
	if err != nil {
		t.Fatalf("signing failed: %v", err)
	}
	if !MLDSAVerify(key.PublicKey, msg, sig) {
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
		if !MLDSAVerify(key.PublicKey, msg, sig) {
			t.Error("signature should verify for fuzzed input")
		}
	})
}

func BenchmarkMLDSASign(b *testing.B) {
	key, _ := GenerateMLDSAKeyPair()
	msg := []byte("benchmark message")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MLDSASign(key.PrivateKey, msg)
	}
}

func BenchmarkMLDSAVerify(b *testing.B) {
	key, _ := GenerateMLDSAKeyPair()
	msg := []byte("benchmark message")
	sig, _ := MLDSASign(key.PrivateKey, msg)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MLDSAVerify(key.PublicKey, msg, sig)
	}
}

func TestMLDSASignWithInvalidKey(t *testing.T) {
	_, err := MLDSASign([]byte{}, []byte("msg"))
	if err == nil {
		t.Error("expected error for empty private key")
	}
}

func TestMLDSAVerifyWithInvalidKey(t *testing.T) {
	key, _ := GenerateMLDSAKeyPair()
	msg := []byte("msg")
	sig, _ := MLDSASign(key.PrivateKey, msg)
	invalidPub := []byte{}
	if MLDSAVerify(invalidPub, msg, sig) {
		t.Error("verify should fail with invalid public key")
	}
}

func TestMLDSAVerifyWithTamperedSignature(t *testing.T) {
	key, _ := GenerateMLDSAKeyPair()
	msg := []byte("msg")
	sig, _ := MLDSASign(key.PrivateKey, msg)
	if len(sig) > 0 {
		sig[0] ^= 0xFF // tamper signature
	}
	if MLDSAVerify(key.PublicKey, msg, sig) {
		t.Error("verify should fail with tampered signature")
	}
}
