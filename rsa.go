package gocrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

type GoRSACrypto struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewRSAGoCrypto
func NewRSACrypto() (*GoRSACrypto, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return &GoRSACrypto{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}, nil
}

// PrivateKeyToPEM
func (g *GoRSACrypto) PrivateKeyToPEM() ([]byte, error) {
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(g.privateKey),
	}
	return pem.EncodeToMemory(block), nil
}

// PublicKeyToPEM
func (g *GoRSACrypto) PublicKeyToPEM() ([]byte, error) {
	pubASN1, err := x509.MarshalPKIXPublicKey(g.publicKey)
	if err != nil {
		return nil, err
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	}
	return pem.EncodeToMemory(block), nil
}

// Encrypt
func (g *GoRSACrypto) Encrypt(plaintext string) (string, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, g.publicKey, []byte(plaintext), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt
func (g *GoRSACrypto) Decrypt(ciphertext string) (string, error) {

	// Decode the base64
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext: %v", err)
	}

	plaintextBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, g.privateKey, ciphertextBytes, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt ciphertext: %v", err)
	}

	plaintext := string(plaintextBytes)
	return plaintext, nil
}
