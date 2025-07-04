package pq

import (
	"testing"
)

func TestGenerateMLKEMKeyPair(t *testing.T) {
	key, err := GenerateMLKEMKeyPair()
	if err != nil {
		t.Fatalf("failed to generate ML-KEM key pair: %v", err)
	}
	if len(key.PublicKey) == 0 || len(key.PrivateKey) == 0 {
		t.Error("key pair should not be empty")
	}
}

func TestMLKEMEncapsulateAndDecapsulate(t *testing.T) {
	key, _ := GenerateMLKEMKeyPair()
	shared, ct, err := MLKEMEncapsulate(key.PublicKey)
	if err != nil {
		t.Fatalf("encapsulation failed: %v", err)
	}
	shared2, err := MLKEMDecapsulate(key.PrivateKey, ct)
	if err != nil {
		t.Fatalf("decapsulation failed: %v", err)
	}
	if len(shared) != len(shared2) {
		t.Error("shared secrets should have same length")
	}
}

func FuzzMLKEMEncapsulateAndDecapsulate(f *testing.F) {
	key, _ := GenerateMLKEMKeyPair()
	f.Fuzz(func(t *testing.T, msg []byte) {
		_, ct, err := MLKEMEncapsulate(key.PublicKey)
		if err != nil {
			t.Skip()
		}
		_, err = MLKEMDecapsulate(key.PrivateKey, ct)
		if err != nil {
			t.Error("decapsulation should not fail for fuzzed input")
		}
	})
}

func BenchmarkMLKEMEncapsulate(b *testing.B) {
	key, _ := GenerateMLKEMKeyPair()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MLKEMEncapsulate(key.PublicKey)
	}
}

func BenchmarkMLKEMDecapsulate(b *testing.B) {
	key, _ := GenerateMLKEMKeyPair()
	_, ct, _ := MLKEMEncapsulate(key.PublicKey)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MLKEMDecapsulate(key.PrivateKey, ct)
	}
}
