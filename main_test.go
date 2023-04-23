package gocrypto_test

import (
	"crypto/rand"
	"testing"

	"github.com/pcpratheesh/gocrypto"
)

func GenerateRandomKey() ([]byte, error) {
	// Generate a 256-bit random key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return key, err
}

func TestEncryptDecrypt(t *testing.T) {
	key := []byte("my-secret-key-16")
	g := gocrypto.NewAESCrypto(key)

	plaintext := "my sensitive information"
	ciphertext, err := g.Encrypt(plaintext)
	if err != nil {
		t.Errorf("Unexpected error during encryption: %v", err)
	}

	if ciphertext == plaintext {
		t.Errorf("Encryption failed, ciphertext is equal to plaintext")
	}

	decrypted, err := g.Decrypt(ciphertext)
	if err != nil {
		t.Errorf("Unexpected error during decryption: %v", err)
	}

	if decrypted != plaintext {
		t.Errorf("Decryption failed, expected %q but got %q", plaintext, decrypted)
	}
}

func TestInvalidCiphertext(t *testing.T) {
	key, _ := GenerateRandomKey()
	g := gocrypto.NewAESCrypto(key)

	_, err := g.Decrypt("invalid-ciphertext")
	if err == nil {
		t.Errorf("Expected error due to invalid ciphertext, but got none")
	}
}

func TestShortCiphertext(t *testing.T) {
	key := []byte("my-secret-key")
	g := gocrypto.NewAESCrypto(key)

	shortCiphertext := "short-ciphertext"
	_, err := g.Decrypt(shortCiphertext)
	if err == nil {
		t.Errorf("Expected error due to short ciphertext, but got none")
	}
}
