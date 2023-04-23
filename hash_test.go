package gocrypto

import (
	"testing"
)

func TestGoHashCrypto(t *testing.T) {
	// Create a new GoHashCrypto instance
	g, err := NewHashCrypto()
	if err != nil {
		t.Errorf("failed to create GoHashCrypto instance: %v", err)
	}

	// Test SetHashFunction method
	err = g.SetHashFunction("SHA512")
	if err != nil {
		t.Errorf("SetHashFunction failed: %v", err)
	}

	// Test HashString method
	hashedData, err := g.HashString("test data")
	if err != nil {
		t.Errorf("HashString failed: %v", err)
	}
	expectedHash := "Dh4h7PEF7IU9JNcohnrXBhPCFmOkaTB0sqNhnBvTnWa1iMM3I7tGbHJCToDjymPCSQeKs0e6uUKFAOfuQwWdDQ=="
	if hashedData != expectedHash {
		t.Errorf("HashString returned incorrect hash: got %s, want %s", hashedData, expectedHash)
	}

	// Test CheckPasswordHash method
	password := "password123"
	hashedPassword, _ := g.HashString(password)
	result := g.CheckPasswordHash(password, hashedPassword)
	if !result {
		t.Errorf("CheckPasswordHash failed: password and hashed password should match")
	}

	// Test invalid hash function
	err = g.SetHashFunction("invalid hash function")
	if err == nil {
		t.Errorf("SetHashFunction should return error on nonexistent hash function")
	}
}

func TestCheckSignature(t *testing.T) {
	// Generate a new GoHashCrypto instance with randomly generated RSA keys
	g, err := NewHashCrypto()
	if err != nil {
		t.Fatalf("Failed to generate GoHashCrypto instance: %v", err)
	}

	// Define a test password and hash value
	password := "password123"
	hashedPassword, _ := g.HashString(password)

	// Verify signature using public key
	result := g.CheckSignature(hashedPassword)
	if !result {
		t.Errorf("CheckSignature failed: signature is invalid")
	}
}
