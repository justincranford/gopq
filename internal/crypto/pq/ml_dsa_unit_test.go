package pq

import (
	"testing"
)

func logTestStartEnd(tb testing.TB) {
	tb.Logf("START: %s", tb.Name())
	tb.Cleanup(func() {
		tb.Logf("END: %s", tb.Name())
	})
}

func init() {
	testing.Init()
}

func TestMain(m *testing.M) {
	// Optionally, add global setup/teardown here.
	m.Run()
}

// Add logTestStartEnd to every test/fuzz/benchmark function above:
func init() {
	// This is a placeholder to ensure the helper is included.
}

// --- Add logTestStartEnd to each test function ---

// For TestGenerateMLDSAKeyPair:
func TestGenerateMLDSAKeyPair(t *testing.T) {
	logTestStartEnd(t)
	key, err := GenerateMLDSAKeyPair()
	if err != nil {
		t.Fatalf("failed to generate ML-DSA key pair: %v", err)
	}
	if len(key.PublicKey) == 0 || len(key.PrivateKey) == 0 {
		t.Error("key pair should not be empty")
	}
}

// For TestMLDSASignAndVerify:
func TestMLDSASignAndVerify(t *testing.T) {
	logTestStartEnd(t)
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

// For TestMLDSASignWithInvalidKey:
func TestMLDSASignWithInvalidKey(t *testing.T) {
	logTestStartEnd(t)
	_, err := MLDSASign([]byte{}, []byte("msg"))
	if err == nil {
		t.Error("expected error for empty private key")
	}
}

// For TestMLDSAVerifyWithInvalidKey:
func TestMLDSAVerifyWithInvalidKey(t *testing.T) {
	logTestStartEnd(t)
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
	logTestStartEnd(t)
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
	if !isVerify {
		t.Error("verify should succeed for original signature")
	}
	if err != nil {
		t.Fatalf("verifying failed: %v", err)
	}
	if len(sig) > 0 {
		sig[0] ^= 0xFF // tamper signature
	}
	isVerify, err = MLDSAVerify(key.PublicKey, msg, sig)
	if isVerify {
		t.Error("verify should fail with tampered signature")
	}
	if err != nil {
		t.Fatalf("verifying failed: %v", err)
	}
}
