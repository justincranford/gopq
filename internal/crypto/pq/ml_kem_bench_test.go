package pq

import "testing"

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
