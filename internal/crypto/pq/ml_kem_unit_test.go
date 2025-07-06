package pq

import (
	"crypto/subtle"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateMLKEMKeyPair(t *testing.T) {
	keyPair, err := GenerateMLKEMKeyPair()
	require.NoError(t, err, "failed to generate ML-KEM key pair")
	require.NotNil(t, keyPair.PublicKey, "public key should not be nil")
	require.NotNil(t, keyPair.PrivateKey, "private key should not be nil")

	publicKeyBytes, err := MarshalPublicKey(keyPair.PublicKey)
	require.NoError(t, err, "marshal public key failed")
	require.NotEmpty(t, publicKeyBytes, "public key bytes should not be empty")

	privateKeyBytes, err := MarshalPrivateKey(keyPair.PrivateKey)
	require.NoError(t, err, "marshal private key failed")
	require.NotEmpty(t, publicKeyBytes, "private key bytes should not be empty")

	publicKey, err := UnmarshalPublicKey(publicKeyBytes)
	require.NoError(t, err, "failed to unmarshall public key")
	require.NotNil(t, publicKey, "public key should not be nil")

	privateKey, err := UnmarshalPrivateKey(privateKeyBytes)
	require.NoError(t, err, "failed to unmarshall private key")
	require.NotNil(t, privateKey, "private key should not be nil")
}

func TestMLKEMEncapsulateAndDecapsulate(t *testing.T) {
	keyPair, err := GenerateMLKEMKeyPair()
	require.NoError(t, err, "failed to generate ML-KEM key pair")
	require.NotNil(t, keyPair.PublicKey, "public key should not be nil")
	require.NotNil(t, keyPair.PrivateKey, "private key should not be nil")

	ciphertext, sharedSecret, err := MLKEMEncapsulate(keyPair.PublicKey)
	require.NoError(t, err, "failed to encapsulate public key")
	require.NotEmpty(t, ciphertext, "ciphertext should not be empty")
	require.NotEmpty(t, sharedSecret, "sharedSecret should not be empty")

	sharedSecret2, err := MLKEMDecapsulate(keyPair.PrivateKey, ciphertext)
	require.NoError(t, err, "failed to decapsulate ciphertext with private key")
	require.NotEmpty(t, sharedSecret2, "sharedSecret2 should not be empty")

	if subtle.ConstantTimeCompare(sharedSecret, sharedSecret2) == 0 {
		t.Error("shared secrets were expected to match, decapculate with private key and ciphertext should have returned a same shared secret as encapculate with public key")
	}
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
	require.NotNil(t, keyPair.PublicKey, "public key should not be nil")
	require.NotNil(t, keyPair.PrivateKey, "private key should not be nil")

	ciphertext, sharedSecret, err := MLKEMEncapsulate(keyPair.PublicKey)
	require.NoError(t, err, "failed to encapsulate public key")
	require.NotEmpty(t, ciphertext, "ciphertext should not be empty")
	require.NotEmpty(t, sharedSecret, "sharedSecret should not be empty")

	ciphertext[0] ^= 0xFF // tamper ciphertext

	sharedSecretTampered, err := MLKEMDecapsulate(keyPair.PrivateKey, ciphertext)
	require.NoError(t, err, "failed to decapsulate ciphertext with private key")
	require.NotEmpty(t, sharedSecretTampered, "sharedSecretTampered should not be empty")

	if subtle.ConstantTimeCompare(sharedSecret, sharedSecretTampered) == 1 {
		t.Error("shared secrets were expected to be different, decapsulated with tampered ciphertext should have returned a different shared secret")
	}
}

func TestMLKEMDecapsulateWithWrongKey(t *testing.T) {
	keyPair1, err := GenerateMLKEMKeyPair()
	require.NoError(t, err, "failed to generate ML-KEM key pair")
	require.NotNil(t, keyPair1.PublicKey, "public key 1 should not be nil")
	require.NotNil(t, keyPair1.PrivateKey, "private key 1 should not be nil")

	keyPair2, err := GenerateMLKEMKeyPair()
	require.NoError(t, err, "failed to generate ML-KEM key pair")
	require.NotNil(t, keyPair2.PublicKey, "public key 2 should not be nil")
	require.NotNil(t, keyPair2.PrivateKey, "private key 2 should not be nil")

	ciphertext, sharedSecret, err := MLKEMEncapsulate(keyPair1.PublicKey)
	require.NoError(t, err, "failed to encapsulate public key 1")
	require.NotEmpty(t, ciphertext, "ciphertext should not be empty")
	require.NotEmpty(t, sharedSecret, "sharedSecret should not be empty")

	sharedSecretWrong, err := MLKEMDecapsulate(keyPair2.PrivateKey, ciphertext)
	require.NoError(t, err, "failed to decapsulate ciphertext with private key 2")
	require.NotEmpty(t, sharedSecretWrong, "sharedSecretWrong should not be empty")

	if subtle.ConstantTimeCompare(sharedSecret, sharedSecretWrong) == 1 {
		t.Error("shared secrets were expected to be different, decapsulated with wrong private key should have returned a different shared secret")
	}
}
