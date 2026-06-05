# crypto

A simple Go library for AES-GCM symmetric encryption and decryption.

## Features

- **AES-GCM encryption**: Authenticated encryption providing both confidentiality and integrity
- **Random nonce generation**: Each encryption produces a unique ciphertext
- **Hex encoding**: Ciphertexts are hex-encoded for easy storage and transmission
- **Key size validation**: Supports AES-128 (16 bytes), AES-192 (24 bytes), and AES-256 (32 bytes)

## Installation

```bash
go get github.com/chrix75/rypto
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/chrix75/rypto"
)

func main() {
	// Set the secret key (AES-256: 32 bytes)
	crypto.SecretKey = []byte("8nFMVCEPqgiwy3ApkL1IqnqkcEJlVmRt")

	// Encrypt a message
	plaintext := "Hello, World!"
	encrypted, err := crypto.Encrypt(plaintext)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted:", encrypted)

	// Decrypt the message
	decrypted, err := crypto.Decrypt(encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted:", decrypted)
}
```

## API Reference

### Variables

```go
var SecretKey []byte
```
The symmetric key used for AES encryption and decryption. Must be set before calling `Encrypt` or `Decrypt`. Valid key sizes are 16 (AES-128), 24 (AES-192), or 32 (AES-256) bytes.

### Functions

```go
func Encrypt(plaintext string) (string, error)
```
Encrypts the given plaintext using AES-GCM with the configured `SecretKey`. Returns the ciphertext as a hex-encoded string. Returns an error if `SecretKey` is not set or has an invalid size.

```go
func Decrypt(ciphertextHex string) (string, error)
```
Decrypts the given hex-encoded ciphertext using AES-GCM with the configured `SecretKey`. Returns the original plaintext as a string. Returns an error if `SecretKey` is not set, has an invalid size, or if the ciphertext is corrupted.

## Security Considerations

- **Key management**: The security of this library depends on the secrecy of `SecretKey`. Never hardcode keys in source code or commit them to version control.
- **Key rotation**: Regularly rotate encryption keys according to your security policy.
- **Key size**: Always use the largest key size practical for your use case (preferably AES-256 with 32-byte keys).
- **Randomness**: The library uses `crypto/rand` for nonce generation, which is cryptographically secure.

## Examples

### Encrypting and Decrypting Strings

```go
crypto.SecretKey = []byte("my-32-byte-long-secret-key-12345")

encrypted, _ := crypto.Encrypt("sensitive data")
decrypted, _ := crypto.Decrypt(encrypted)
// decrypted == "sensitive data"
```

### Handling Errors

```go
crypto.SecretKey = []byte("invalid-key-size")

_, err := crypto.Encrypt("test")
// err != nil (invalid key size)

crypto.SecretKey = nil
_, err = crypto.Encrypt("test")
// err != nil (SecretKey not set)
```

## Testing

```bash
go test ./...
```

## License

MIT License

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.
