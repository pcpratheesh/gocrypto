package example

import (
	"encoding/pem"
	"fmt"

	"github.com/pcpratheesh/gocrypto"
)

func ExampleRSA() {
	// Create a new GoRSACrypto instance with randomly generated RSA keys
	g, err := gocrypto.NewRSACrypto()
	if err != nil {
		fmt.Printf("Failed to create GoRSACrypto instance: %v", err)
		return
	}

	// Convert private key to PEM format
	privateKeyPEM, err := g.PrivateKeyToPEM()
	if err != nil {
		fmt.Printf("Failed to convert private key to PEM format: %v", err)
		return
	}
	fmt.Println(string(privateKeyPEM))

	// Convert public key to PEM format
	publicKeyPEM, err := g.PublicKeyToPEM()
	if err != nil {
		fmt.Printf("Failed to convert public key to PEM format: %v", err)
		return
	}
	fmt.Println(string(publicKeyPEM))

	// Encrypt plaintext using public key
	plaintext := "Hello, world!"
	ciphertext, err := g.Encrypt(plaintext)
	if err != nil {
		fmt.Printf("Failed to encrypt plaintext: %v", err)
		return
	}
	fmt.Println(ciphertext)

	// Decrypt ciphertext using private key
	decryptedText, err := g.Decrypt(ciphertext)
	if err != nil {
		fmt.Printf("Failed to decrypt ciphertext: %v", err)
		return
	}
	fmt.Println(decryptedText)

	// Decode the private key PEM file
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		fmt.Println("Failed to decode private key PEM file")
		return
	}

}
