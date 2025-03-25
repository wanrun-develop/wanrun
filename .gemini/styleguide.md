# Company X Golang Style Guide

# Introduction
This style guide outlines the coding conventions for Golang code developed at Company X.
It's based on Go’s standard conventions and best practices, with some modifications to address specific needs and preferences within our organization.

# Key Principles
* **Readability:** Code should be easy to understand for all team members.
* **Maintainability:** Code should be easy to modify and extend.
* **Consistency:** Adhering to a consistent style across all projects improves collaboration and reduces errors.
* **Performance:** Go emphasizes performance and simplicity. Code should be efficient and avoid unnecessary complexity.
* **Idiomatic Go:** Follow Go’s conventions, avoiding over-engineering or patterns that diverge from standard Go practices.

# Code Formatting
* **Use `gofmt`:** All code must be formatted using `gofmt`. This ensures consistency and readability.
* **Line Length:** No hard line length limit, but keep lines reasonably short for readability.
* **Indentation:** Use tabs for indentation (Go’s standard). Align with spaces where necessary.

# Imports
* **Group imports:**
  * Standard library imports
  * Third-party package imports
  * Internal packages
* **Import order:** Separate groups with blank lines and sort imports alphabetically within each group.

# Naming Conventions
* **Variables:** Use camelCase for variables (e.g., `userName`, `totalCount`).
* **Constants:** Use PascalCase for exported constants and camelCase for unexported ones (e.g., `MaxValue`, `databaseName`).
* **Functions:** Use PascalCase for exported functions and camelCase for unexported ones (e.g., `CalculateTotal()`, `processData()`).
* **Packages:** Use short, lowercase names (e.g., `userutil`, `paymentgateway`). Avoid underscores or mixed case.
* **Types and Structs:** Use PascalCase for exported types (e.g., `UserManager`, `PaymentProcessor`).

# Comments
* **Use complete sentences:** Comments should be grammatically correct and end with a period.
* **Documentation Comments:** For exported items, provide comments directly above the declaration (e.g., `// CalculateTotal sums all the values and returns the result.`).
* **Inline Comments:** Avoid unless necessary to explain complex logic.

# Logging
* **Use a standard logging framework:** Prefer using `log` or `zap` for structured logging.
* **Log at appropriate levels:** INFO, DEBUG, WARNING, ERROR, FATAL.
* **Include context:** Include relevant information to aid debugging.

# Error Handling
* **Use `error` values:** Functions that can fail should return an `error` type as the last return value.
* **Avoid panic:** Use `panic` only for unrecoverable errors or programming errors.
* **Use `fmt.Errorf()` or `errors.New()`:** Provide informative error messages.

# Tooling
* **Code Formatter:** `gofmt` - Enforces consistent formatting automatically.
* **Linter:** `golint`, `staticcheck`, `errcheck` - Identifies potential issues and style violations.
* **Testing:** Use `go test` for writing and running unit tests.

# Example
```go
// Package userauth provides utilities for user authentication.

package userauth

import (
    "crypto/sha256"
    "encoding/hex"
    "errors"
    "log"
)

// HashPassword hashes a password using SHA-256.
func HashPassword(password string) (string, error) {
    if password == "" {
        return "", errors.New("password cannot be empty")
    }

    hash := sha256.New()
    hash.Write([]byte(password))
    return hex.EncodeToString(hash.Sum(nil)), nil
}

// AuthenticateUser checks if the provided password matches the stored hash.
func AuthenticateUser(storedHash, password string) bool {
    if storedHash == "" || password == "" {
        log.Println("Invalid input for authentication")
        return false
    }

    hash, err := HashPassword(password)
    if err != nil {
        log.Println("Error generating hash:", err)
        return false
    }

    return storedHash == hash
}
```

# LLM Answer
* **Answer format:** Always respond in Japanese.
