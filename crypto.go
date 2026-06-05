// Package crypto provides simple AES-GCM encryption and decryption utilities.
// It uses the standard Go crypto libraries to perform authenticated encryption,
// ensuring both confidentiality and integrity of the encrypted data.
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

// SecretKey is the symmetric key used for AES encryption and decryption.
// It must be set before calling Encrypt or Decrypt functions.
// Valid key sizes are 16 (AES-128), 24 (AES-192), or 32 (AES-256) bytes.
var SecretKey []byte

// Encrypt encrypts the given plaintext using AES-GCM with the configured SecretKey.
// It generates a random nonce for each encryption, ensuring that encrypting the
// same plaintext multiple times produces different ciphertexts.
// Returns the ciphertext as a hex-encoded string.
// Returns an error if SecretKey is not set or has an invalid size.
func Encrypt(plaintext string) (string, error) {
	if len(SecretKey) == 0 {
		return "", errors.New("SecretKey must be set before encryption")
	}

	block, err := aes.NewCipher(SecretKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the given hex-encoded ciphertext using AES-GCM with the configured SecretKey.
// The nonce is extracted from the ciphertext (prepended during encryption).
// Returns the original plaintext as a string.
// Returns an error if SecretKey is not set, has an invalid size, or if the ciphertext is corrupted.
func Decrypt(ciphertextHex string) (string, error) {
	if len(SecretKey) == 0 {
		return "", errors.New("SecretKey must be set before decryption")
	}

	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(SecretKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
