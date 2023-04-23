package gocrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// Model
type GoCrypto struct {
	key []byte
}

// NewGoCrypto
func NewAESCrypto(key []byte) *GoCrypto {
	return &GoCrypto{
		key: key,
	}
}

// Encrypt
func (g *GoCrypto) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(g.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt
func (g *GoCrypto) Decrypt(ciphertext string) (string, error) {
	block, err := aes.NewCipher(g.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(decodedCiphertext) < gcm.NonceSize() {
		return "", errors.New("ciphertext is too short")
	}

	nonce := decodedCiphertext[:gcm.NonceSize()]
	decodedCiphertext = decodedCiphertext[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, decodedCiphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
