package pq

import "testing"

func BenchmarkMLDSASign(b *testing.B) {
	logTestStartEnd(b)
	key, err := GenerateMLDSAKeyPair()
	if err != nil {
		b.Fatalf("failed to generate ML-DSA key pair: %v", err)
	}
	message := []byte("benchmark message")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, signErr := MLDSASign(key.PrivateKey, message)
		if signErr != nil {
			b.Fatalf("signing failed: %v", signErr)
		}
	}
}

func BenchmarkMLDSAVerify(b *testing.B) {
	logTestStartEnd(b)
	key, err := GenerateMLDSAKeyPair()
	if err != nil {
		b.Fatalf("failed to generate ML-DSA key pair: %v", err)
	}
	message := []byte("benchmark message")
	signature, signErr := MLDSASign(key.PrivateKey, message)
	if signErr != nil {
		b.Fatalf("signing failed: %v", signErr)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, verifyErr := MLDSAVerify(key.PublicKey, message, signature)
		if verifyErr != nil {
			b.Fatalf("verifying failed: %v", verifyErr)
		}
	}
}
