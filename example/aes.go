package example

import (
	"fmt"

	"github.com/pcpratheesh/gocrypto"
)

func ExampleAES() {
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
