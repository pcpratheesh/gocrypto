package gocrypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
)

type GoHashCrypto struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	hashFunc   crypto.Hash
}

// NewHashCrypto
func NewHashCrypto() (*GoHashCrypto, error) {
	keyPair, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate RSA key pair: %v", err)
	}
	g := &GoHashCrypto{
		privateKey: keyPair,
		publicKey:  &keyPair.PublicKey,
		hashFunc:   crypto.SHA256,
	}

	return g, nil
}

// SetHashFunction
func (g *GoHashCrypto) SetHashFunction(hashFunc string) error {
	switch hashFunc {
	case "MD5":
		g.hashFunc = crypto.MD5
	case "SHA256":
		g.hashFunc = crypto.SHA256
	case "SHA512":
		g.hashFunc = crypto.SHA512
	default:
		return errors.New("invalid hash function")
	}
	return nil
}

// HashString
// it will return the hash of given string using the hash function used
func (g *GoHashCrypto) HashString(data string) (string, error) {
	hasher := g.hashFunc.New()
	_, err := hasher.Write([]byte(data))
	if err != nil {
		return "", fmt.Errorf("error writing data to hash function: %v", err)
	}

	sum := hasher.Sum(nil)
	hashedData := base64.StdEncoding.EncodeToString(sum)
	return hashedData, nil
}

// CheckSignature
func (g *GoHashCrypto) CheckSignature(hashedPassword string) bool {
	if g.publicKey == nil {
		return false
	}
	decodedHash, _ := base64.StdEncoding.DecodeString(hashedPassword)
	signature, _ := rsa.SignPKCS1v15(rand.Reader, g.privateKey, g.hashFunc, decodedHash)

	// Verify signature using public key
	err := rsa.VerifyPKCS1v15(g.publicKey, g.hashFunc, decodedHash, signature)
	return err == nil
}

// CheckPasswordHash
// checks if a given password matches the hashed password
func (g *GoHashCrypto) CheckPasswordHash(password, hashedPassword string) bool {
	hash, err := g.HashString(password)
	if err != nil {
		return false
	}

	return hash == hashedPassword
}
