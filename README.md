# gocrypto
A package that provides a cryptographic algorithms and protocols in Go, including support for encryption, decryption, signing, and verification

This GoCrypto package provides a simple and secure way to encrypt and decrypt sensitive information in a Go program using the AES-GCM encryption algorithm. It allows developers to protect confidential data while it is stored or transmitted over insecure channels, which helps to prevent unauthorized access, tampering, or eavesdropping. This package can be useful in various scenarios such as securing passwords, credit card numbers, personal identification information, and other sensitive data that need to be protected from prying eyes.

### Usage

1. Import the package:

```go
import "github.com/example/gocrypto"
```

2. Create an instance of GoCrypto with an encryption key:

```go
key := []byte("my-secret-key")
g := gocrypto.NewGoCrypto(key)
```

3. Encrypt a plaintext message:

```go
plaintext := "my sensitive information"
ciphertext, err := g.Encrypt(plaintext)
if err != nil {
  // handle error
}
```

4. Decrypt a ciphertext message:

```go
ciphertext := "encrypted message"
plaintext, err := g.Decrypt(ciphertext)
if err != nil {
  // handle error
}
```


## Supporting Encryptions algorithms
- [RSA](example/rsa.md)
- [AES](example/aes.md)
- [HASH](example/hash.md)

### License

This package is licensed under the MIT License. See the LICENSE file for more details.
