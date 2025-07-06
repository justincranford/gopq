package pq

import (
	"crypto/subtle"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMarshallUnmarshallMLKEMKeyPair(t *testing.T) {
	keyPair, err := GenerateMLKEMKeyPair()
	require.NoError(t, err, "failed to generate ML-KEM key pair")
	require.NotNil(t, keyPair, "expected non-nil key pair")
	require.NotNil(t, keyPair.PublicKey, "expected non-nil public key")
	require.NotNil(t, keyPair.PrivateKey, "expected non-nil private key")

	publicKeyBytes, err := MarshalPublicKey(keyPair.PublicKey)
	require.NoError(t, err, "failed to marshal public key")
	require.NotEmpty(t, publicKeyBytes, "expected non-empty public key bytes")

	privateKeyBytes, err := MarshalPrivateKey(keyPair.PrivateKey)
	require.NoError(t, err, "failed to marshal private key")
	require.NotEmpty(t, publicKeyBytes, "expected non-empty private key bytes")

	publicKey, err := UnmarshalPublicKey(publicKeyBytes)
	require.NoError(t, err, "failed to unmarshall public key")
	require.NotNil(t, keyPair.PublicKey, "expected non-nil public key")
	require.True(t, publicKey.Equal(keyPair.PublicKey), "expected unmarshalled public key to match the original public key")

	privateKey, err := UnmarshalPrivateKey(privateKeyBytes)
	require.NoError(t, err, "failed to unmarshall private key")
	require.NotNil(t, keyPair.PrivateKey, "expected non-nil private key")
	require.True(t, privateKey.Equal(keyPair.PrivateKey), "expected unmarshalled private key to match the original private key")
}

func TestMLKEMEncapsulateAndDecapsulate(t *testing.T) {
	keyPair, err := GenerateMLKEMKeyPair()
	require.NoError(t, err, "failed to generate ML-KEM key pair")
	require.NotNil(t, keyPair, "expected non-nil key pair")
	require.NotNil(t, keyPair.PublicKey, "expected non-nil public key")
	require.NotNil(t, keyPair.PrivateKey, "expected non-nil private key")

	ciphertext, sharedSecret, err := MLKEMEncapsulate(keyPair.PublicKey)
	require.NoError(t, err, "failed to encapsulate public key")
	require.NotEmpty(t, ciphertext, "expected non-empty ciphertext")
	require.NotEmpty(t, sharedSecret, "expected non-empty sharedSecret")

	sharedSecret2, err := MLKEMDecapsulate(keyPair.PrivateKey, ciphertext)
	require.NoError(t, err, "failed to decapsulate ciphertext with private key")
	require.NotEmpty(t, sharedSecret2, "expected non-empty sharedSecret2")

	require.NotZero(t, subtle.ConstantTimeCompare(sharedSecret, sharedSecret2), "expected shared secrets to match")
}

func TestMLKEMEncapsulateWithInvalidKey(t *testing.T) {
	ciphertext, sharedSecret, err := MLKEMEncapsulate(nil)
	require.Error(t, err, "expected error for encapsulate with nil public key")
	require.Nil(t, ciphertext, "expected nil ciphertext for encapsulate with nil public key")
	require.Nil(t, sharedSecret, "expected nil sharedSecret for encapsulate with nil public key")
}

func TestMLKEMDecapsulateWithInvalidKey(t *testing.T) {
	ciphertext, sharedSecret, err := MLKEMEncapsulate(nil)
	require.Error(t, err, "expected error for encapsulate with nil public key")
	require.Nil(t, ciphertext, "expected nil ciphertext for encapsulate with nil public key")
	require.Nil(t, sharedSecret, "expected nil sharedSecret for encapsulate with nil public key")

	sharedSecret2, err := MLKEMDecapsulate(nil, ciphertext)
	require.Error(t, err, "expected error for decapsulate ciphertext with nil private key")
	require.Nil(t, sharedSecret2, "expected nil sharedSecret2 for decapsulate with nil private key")
}

func TestMLKEMDecapsulateWithTamperedCiphertext(t *testing.T) {
	keyPair, err := GenerateMLKEMKeyPair()
	require.NoError(t, err, "failed to generate ML-KEM key pair")
	require.NotNil(t, keyPair, "expected non-nil key pair")
	require.NotNil(t, keyPair.PublicKey, "expected non-nil public key")
	require.NotNil(t, keyPair.PrivateKey, "expected non-nil private key")

	ciphertext, sharedSecret, err := MLKEMEncapsulate(keyPair.PublicKey)
	require.NoError(t, err, "failed to encapsulate public key")
	require.NotEmpty(t, ciphertext, "expected non-empty ciphertext")
	require.NotEmpty(t, sharedSecret, "expected non-empty sharedSecret")

	ciphertext[0] ^= 0xFF // tamper ciphertext

	sharedSecretTampered, err := MLKEMDecapsulate(keyPair.PrivateKey, ciphertext)
	require.NoError(t, err, "failed to decapsulate ciphertext with private key")
	require.NotEmpty(t, sharedSecretTampered, "expected non-empty sharedSecretTampered")

	require.Zero(t, subtle.ConstantTimeCompare(sharedSecret, sharedSecretTampered), "expected different shared secret for decapsulate with tampered ciphertext")
}

func TestMLKEMDecapsulateWithWrongKey(t *testing.T) {
	keyPair1, err := GenerateMLKEMKeyPair()
	require.NoError(t, err, "failed to generate ML-KEM key pair 1")
	require.NotNil(t, keyPair1, "expected non-nil key pair 1")
	require.NotNil(t, keyPair1.PublicKey, "expected non-nil public key 1")
	require.NotNil(t, keyPair1.PrivateKey, "expected non-nil private key 1")

	keyPair2, err := GenerateMLKEMKeyPair()
	require.NoError(t, err, "failed to generate ML-KEM key pair 2")
	require.NotNil(t, keyPair2, "expected non-nil key pair 2")
	require.NotNil(t, keyPair2.PublicKey, "expected non-nil public key 2")
	require.NotNil(t, keyPair2.PrivateKey, "expected non-nil private key 2")

	ciphertext, sharedSecret, err := MLKEMEncapsulate(keyPair1.PublicKey)
	require.NoError(t, err, "failed to encapsulate public key 1")
	require.NotEmpty(t, ciphertext, "expected non-empty ciphertext")
	require.NotEmpty(t, sharedSecret, "expected non-empty sharedSecret")

	sharedSecretWrong, err := MLKEMDecapsulate(keyPair2.PrivateKey, ciphertext)
	require.NoError(t, err, "failed to decapsulate ciphertext with private key 2")
	require.NotEmpty(t, sharedSecretWrong, "expected non-empty sharedSecretWrong")

	require.Zero(t, subtle.ConstantTimeCompare(sharedSecret, sharedSecretWrong), "expected different shared secret for decapsulate with wrong private key")
}
