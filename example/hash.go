package example

import (
	"fmt"

	"github.com/pcpratheesh/gocrypto"
)

func ExampleHASH() {
	// Create a new GoHashCrypto instance with 2048-bit RSA keys
	g, err := gocrypto.NewHashCrypto()
	if err != nil {
		fmt.Printf("Failed to create GoHashCrypto instance: %v", err)
		return
	}

	err = g.SetHashFunction("SHA512")
	if err != nil {
		fmt.Printf("Failed to create SetHashFunction: %v", err)
		return
	}

	// Hash a password and verify it
	password := "password123"
	hashedPassword, _ := g.HashString(password)
	match := g.CheckPasswordHash(password, hashedPassword)
	if match {
		fmt.Println("Password matches!")
	} else {
		fmt.Println("Password does not match.")
	}

	// Checking the signature
	match = g.CheckSignature(hashedPassword)
	if match {
		fmt.Println("Signature is valid!")
	} else {
		fmt.Println("Signature is invalid.")
	}

}
