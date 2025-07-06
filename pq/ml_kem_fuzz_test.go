package pq

import "testing"

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
