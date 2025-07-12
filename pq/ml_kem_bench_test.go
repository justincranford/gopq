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
			ciphertext, sharedSecret, err := MLKEMEncapsulate(keyPair.PublicKey)
			require.NoError(b, err, "Encapsulation should not error")
			require.NotNil(b, ciphertext, "Ciphertext should not be nil")
			require.NotNil(b, sharedSecret, "Shared secret should not be nil")
		}
	})
}

func BenchmarkMLKEMDecapsulate(b *testing.B) {
	keyPair, err := GenerateMLKEMKeyPair()
	require.NoError(b, err, "Key pair generation should not error")

	ciphertext, sharedSecret, err := MLKEMEncapsulate(keyPair.PublicKey)
	require.NoError(b, err, "Encapsulation should not error")
	require.NotNil(b, ciphertext, "Ciphertext should not be nil")
	require.NotNil(b, sharedSecret, "Shared secret should not be nil")

	b.ResetTimer()

	// Use the same ciphertext for all iterations
	b.RunParallel(func(pb *testing.PB) {
		localCiphertext := make([]byte, len(ciphertext))
		copy(localCiphertext, ciphertext)

		for pb.Next() {
			recoveredSecret, err := MLKEMDecapsulate(keyPair.PrivateKey, localCiphertext)
			require.NoError(b, err, "Decapsulation should not error")
			require.NotNil(b, recoveredSecret, "Recovered secret should not be nil")
		}
	})
}
