package pq

import (
	"testing"
)

func TestGenerateMLKEMKeyPair(t *testing.T) {
	key, err := GenerateMLKEMKeyPair()
	if err != nil {
		t.Fatalf("failed to generate ML-KEM key pair: %v", err)
	}
	if key.PublicKey == nil || key.PrivateKey == nil {
		t.Error("key pair should not be nil")
	}
	pubBytes, err := MarshalPublicKey(key.PublicKey)
	if err != nil {
		t.Fatalf("marshal public key failed: %v", err)
	}
	privBytes, err := MarshalPrivateKey(key.PrivateKey)
	if err != nil {
		t.Fatalf("marshal private key failed: %v", err)
	}
	t.Logf("Generated PublicKey: %x", pubBytes)
	t.Logf("Generated PrivateKey: %x", privBytes)
	pub2, err := UnmarshalPublicKey(pubBytes)
	if err != nil {
		t.Fatalf("unmarshal public key failed: %v", err)
	}
	priv2, err := UnmarshalPrivateKey(privBytes)
	if err != nil {
		t.Fatalf("unmarshal private key failed: %v", err)
	}
	if pub2 == nil || priv2 == nil {
		t.Error("unmarshaled keys should not be nil")
	}
}

func TestMLKEMEncapsulateAndDecapsulate(t *testing.T) {
	key, _ := GenerateMLKEMKeyPair()
	pubBytes, _ := MarshalPublicKey(key.PublicKey)
	privBytes, _ := MarshalPrivateKey(key.PrivateKey)
	t.Logf("Test PublicKey: %x", pubBytes)
	t.Logf("Test PrivateKey: %x", privBytes)
	shared, ct, err := MLKEMEncapsulate(key.PublicKey)
	if err != nil {
		t.Fatalf("encapsulation failed: %v", err)
	}
	t.Logf("Encapsulated Ciphertext: %x", ct)
	t.Logf("Encapsulated SharedSecret: %x", shared)
	shared2, err := MLKEMDecapsulate(key.PrivateKey, ct)
	if err != nil {
		t.Fatalf("decapsulation failed: %v", err)
	}
	t.Logf("Decapsulated SharedSecret: %x", shared2)
	if len(shared) != len(shared2) || shared == nil || shared2 == nil {
		t.Error("shared secrets should have same length and not be nil")
	}
	for i := range shared {
		if shared[i] != shared2[i] {
			t.Error("shared secrets do not match")
			break
		}
	}
}

func FuzzMLKEMEncapsulateAndDecapsulate(f *testing.F) {
	key, _ := GenerateMLKEMKeyPair()
	pubBytes, _ := MarshalPublicKey(key.PublicKey)
	privBytes, _ := MarshalPrivateKey(key.PrivateKey)
	f.Fuzz(func(t *testing.T, msg []byte) {
		t.Logf("Fuzz PublicKey: %x", pubBytes)
		t.Logf("Fuzz PrivateKey: %x", privBytes)
		_, ct, err := MLKEMEncapsulate(key.PublicKey)
		if err != nil {
			t.Skip()
		}
		t.Logf("Fuzz Encapsulated Ciphertext: %x", ct)
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
