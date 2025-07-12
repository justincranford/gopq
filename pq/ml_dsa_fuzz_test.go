package pq

import "testing"

func FuzzMLDSASignAndVerify(f *testing.F) {
	key, _ := GenerateMLDSAKeyPair()
	f.Fuzz(func(t *testing.T, msg []byte) {
		signature, signErr := MLDSASign(key.PrivateKey, msg)
		if signErr != nil {
			t.Skip()
		}
		isValid, verifyErr := MLDSAVerify(key.PublicKey, msg, signature)
		if verifyErr != nil {
			t.Fatalf("verifying failed: %v", verifyErr)
		}
		if !isValid {
			t.Error("signature should verify for fuzzed input")
		}
	})
}
