
```go
import (
    "fmt"
    "github.com/pcpratheesh/gocrypto"
)

func main() {
    // Create a new GoHashCrypto instance
    g, err := gocrypto.NewHashCrypto()
    if err != nil {
        fmt.Printf("Failed to create GoHashCrypto instance: %v", err)
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
}
```

## Method Reference

### `NewHashCrypto`
Creates a new `GoHashCrypto` instance with randomly generated RSA keys.

```go
func NewHashCrypto() (*GoHashCrypto, error)
```

Example usage:

```go
g, err := gocrypto.NewHashCrypto()
if err != nil {
    fmt.Printf("Failed to create GoHashCrypto instance: %v", err)
    return
}
```

### `HashString`
Generates a hash value for the given string using the hash function specified in the `GoHashCrypto` instance.

```go
func (g *GoHashCrypto) HashString(data string) (string, error)
```

Example usage:

```go
hashedPassword, err := g.HashString("password123")
if err != nil {
    fmt.Printf("Failed to hash password: %v", err)
    return
}
```

### `CheckPasswordHash`
Checks if the given password matches the previously hashed password.

```go
func (g *GoHashCrypto) CheckPasswordHash(password, hashedPassword string) bool
```

Example usage:

```go
match := g.CheckPasswordHash("password123", hashedPassword)
if match {
    fmt.Println("Password matches!")
} else {
    fmt.Println("Password does not match.")
}
```

### `CheckSignature`
Verifies a digital signature for the given password using RSA digital signatures.

```go
func (g *GoHashCrypto) CheckSignature(hashedPassword string) bool
```

Example usage:

```go
match := g.CheckSignature( hashedPassword)
if match {
    fmt.Println("Signature is valid!")
} else {
    fmt.Println("Signature is invalid.")
}
```
