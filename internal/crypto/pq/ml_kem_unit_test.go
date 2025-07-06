package pq

import "testing"

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
	shared, ct, err := MLKEMEncapsulate(key.PublicKey)
	if err != nil {
		t.Fatalf("encapsulation failed: %v", err)
	}
	shared2, err := MLKEMDecapsulate(key.PrivateKey, ct)
	if err != nil {
		t.Fatalf("decapsulation failed: %v", err)
	}
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

func TestMLKEMEncapsulateWithInvalidKey(t *testing.T) {
	_, _, err := MLKEMEncapsulate(nil)
	if err == nil {
		t.Error("expected error for nil public key")
	}
}

func TestMLKEMDecapsulateWithInvalidKey(t *testing.T) {
	_, ct, _ := MLKEMEncapsulate(nil)
	_, err := MLKEMDecapsulate(nil, ct)
	if err == nil {
		t.Error("expected error for nil private key")
	}
}

func TestMLKEMDecapsulateWithTamperedCiphertext(t *testing.T) {
	key, _ := GenerateMLKEMKeyPair()
	_, ct, _ := MLKEMEncapsulate(key.PublicKey)
	shared2, err := MLKEMDecapsulate(key.PrivateKey, ct)
	if err != nil {
		t.Fatalf("decapsulation failed for original ciphertext: %v", err)
	}
	if len(ct) > 0 {
		ct[0] ^= 0xFF // tamper ciphertext
	}
	sharedTampered, err := MLKEMDecapsulate(key.PrivateKey, ct)
	if err != nil {
		t.Logf("decapsulation failed for tampered ciphertext (acceptable): %v", err)
		return
	}
	if string(shared2) == string(sharedTampered) {
		t.Error("shared secret should differ for tampered ciphertext")
	}
}

func TestMLKEMDecapsulateWithWrongKey(t *testing.T) {
	key1, _ := GenerateMLKEMKeyPair()
	key2, _ := GenerateMLKEMKeyPair()
	shared, ct, _ := MLKEMEncapsulate(key1.PublicKey)
	sharedWrong, err := MLKEMDecapsulate(key2.PrivateKey, ct)
	if err != nil {
		t.Logf("decapsulation failed for wrong private key (acceptable): %v", err)
		return
	}
	if string(shared) == string(sharedWrong) {
		t.Error("shared secret should differ when decapsulated with wrong private key")
	}
}
