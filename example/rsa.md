
# GoRSACrypto

GoRSACrypto is a simple cryptographic library that provides functions for generating RSA keys, converting them to PEM format, and encrypting/decrypting data using RSA encryption.

## Installation

To use the `GoRSACrypto` library in your Go project, you can install it using `go get`:

```go
go get github.com/pcpratheesh/gocrypto
```

## Usage

Here's an example usage of the `GoRSACrypto` package:

```go
import (
    "fmt"
    "github.com/pcpratheesh/gocrypto"
)

func main() {
    // Create a new GoRSACrypto instance with randomly generated RSA keys
    g, err := gocrypto.NewRSACrypto()
    if err != nil {
        fmt.Printf("Failed to create GoRSACrypto instance: %v", err)
        return
    }

    // Encrypt plaintext and decrypt ciphertext
    plaintext := "Hello, world!"
    ciphertext, _ := g.Encrypt(plaintext)
    decryptedText, _ := g.Decrypt(ciphertext)
    if plaintext == decryptedText {
        fmt.Println("Encryption and decryption successful!")
    } else {
        fmt.Println("Encryption and decryption failed.")
    }
}
```



## Method Reference

### `NewRSACrypto`
Creates a new `GoRSACrypto` instance with randomly generated RSA keys.

```go
func NewRSACrypto() (*GoRSACrypto, error)
```

Example usage:

```go
g, err := gocrypto.NewRSACrypto()
if err != nil {
    fmt.Printf("Failed to create GoRSACrypto instance: %v", err)
    return
}
```

### `PrivateKeyToPEM`
Converts the private key to PEM format.

```go
func (g *GoRSACrypto) PrivateKeyToPEM() ([]byte, error)
```

Example usage:

```go
privateKey, err := g.PrivateKeyToPEM()
if err != nil {
    fmt.Printf("Failed to convert private key to PEM format: %v", err)
    return
}
```

### `PublicKeyToPEM`
Converts the public key to PEM format