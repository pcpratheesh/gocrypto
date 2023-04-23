# GoCrypto - AES

GoCrypto is a package for performing encryption and decryption operations using AES encryption.

## Installation

To use GoCrypto in your Go project, simply run:

```bash
go get github.com/pcpratheesh/gocrypto
```

## Usage

Here's an example of how to use GoCrypto to encrypt and decrypt a string:

```go
package main

import (
	"fmt"

	"github.com/pcpratheesh/gocrypto"
)

func main() {
	key := []byte("a very secret key")

	g := gocrypto.NewAESCrypto(key)

	plaintext := "Hello, world!"

	ciphertext, err := g.Encrypt(plaintext)
	if err != nil {
		fmt.Printf("Failed to encrypt plaintext: %v", err)
		return
	}
	fmt.Println(ciphertext)

	decryptedText, err := g.Decrypt(ciphertext)
	if err != nil {
		fmt.Printf("Failed to decrypt ciphertext: %v", err)
		return
	}
	fmt.Println(decryptedText)
}
```

In this example, we create a new `GoCrypto` instance with a given key, then use it to encrypt a plaintext string and print the resulting ciphertext. We then decrypt the ciphertext and print the resulting plaintext.

## Method Reference

### `NewAESCrypto(key []byte) *GoCrypto`

Creates a new `GoCrypto` instance with the given key.

### `(g *GoCrypto) Encrypt(plaintext string) (string, error)`

Encrypts the given plaintext string using AES encryption and returns the resulting base64-encoded ciphertext.

### `(g *GoCrypto) Decrypt(ciphertext string) (string, error)`

Decrypts the given base64-encoded ciphertext using AES decryption and returns the resulting plaintext string.
