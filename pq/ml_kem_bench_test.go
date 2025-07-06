package pq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkMLKEMEncapsulate(b *testing.B) {
	keyPair, err := GenerateMLKEMKeyPair()
	require.NoError(b, err, "Key pair generation should not error")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sharedSecret, ciphertext, err := MLKEMEncapsulate(keyPair.PublicKey)
			require.NoError(b, err, "Encapsulation should not error")
			require.NotNil(b, sharedSecret, "Shared secret should not be nil")
			require.NotNil(b, ciphertext, "Ciphertext should not be nil")
		}
	})
}

func BenchmarkMLKEMDecapsulate(b *testing.B) {
	keyPair, err := GenerateMLKEMKeyPair()
	require.NoError(b, err, "Key pair generation should not error")
	sharedSecret, ciphertext, err := MLKEMEncapsulate(keyPair.PublicKey)
	require.NoError(b, err, "Encapsulation should not error")
	require.NotNil(b, sharedSecret, "Shared secret should not be nil")
	require.NotNil(b, ciphertext, "Ciphertext should not be nil")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			recoveredSecret, err := MLKEMDecapsulate(keyPair.PrivateKey, ciphertext)
			require.NoError(b, err, "Decapsulation should not error")
			require.NotNil(b, recoveredSecret, "Recovered secret should not be nil")
		}
	})
}
