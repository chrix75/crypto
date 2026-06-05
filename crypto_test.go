package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEncryptDecrypt verifies that encryption and decryption work correctly
// with a valid AES-256 key.
func TestEncryptDecrypt(t *testing.T) {
	// given
	SecretKey = []byte("8nFMVCEPqgiwy3ApkL1IqnqkcEJlVmRt")
	text := "Hello world!"

	// when
	encrypted, err := Encrypt(text)
	assert.NoError(t, err)

	decrypted, err := Decrypt(encrypted)
	assert.NoError(t, err)

	// then
	t.Log("Text:" + text)
	t.Log("Encrypted:" + encrypted)
	t.Log("Decrypted:" + decrypted)
	assert.NotEqual(t, text, encrypted)
	assert.Equal(t, text, decrypted)
}

// TestEncryptEmptyKey verifies that encryption fails when SecretKey is not set.
func TestEncryptEmptyKey(t *testing.T) {
	SecretKey = nil
	_, err := Encrypt("test")
	assert.Error(t, err)
}

// TestDecryptEmptyKey verifies that decryption fails when SecretKey is not set.
func TestDecryptEmptyKey(t *testing.T) {
	SecretKey = nil
	_, err := Decrypt("test")
	assert.Error(t, err)
}

// TestDecryptInvalidCiphertext verifies that decryption fails with invalid input.
func TestDecryptInvalidCiphertext(t *testing.T) {
	SecretKey = []byte("8nFMVCEPqgiwy3ApkL1IqnqkcEJlVmRt")
	_, err := Decrypt("invalid-hex-data")
	assert.Error(t, err)
}
