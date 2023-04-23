package gocrypto

import (
	"crypto/rand"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	rsaCrypto, err := NewRSACrypto()
	if err != nil {
		t.Fatalf("Failed to create new RSA crypto: %v", err)
	}

	plaintext := "hello world"

	// encrypt the plaintext with RSA encryption
	ciphertext, err := rsaCrypto.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Failed to encrypt plaintext: %v", err)
	}

	// decrypt the ciphertext with RSA decryption
	decryptedPlaintext, err := rsaCrypto.Decrypt(ciphertext)
	if err != nil {
		t.Fatalf("Failed to decrypt ciphertext: %v", err)
	}

	if decryptedPlaintext != plaintext {
		t.Fatalf("Decrypted plaintext does not match original plaintext")
	}
}

func BenchmarkEncryptDecrypt(b *testing.B) {
	rsaCrypto, _ := NewRSACrypto()
	plaintext := make([]byte, 1024)
	rand.Read(plaintext)

	for i := 0; i < b.N; i++ {
		ciphertext, _ := rsaCrypto.Encrypt(string(plaintext))
		rsaCrypto.Decrypt(ciphertext)
	}
}
