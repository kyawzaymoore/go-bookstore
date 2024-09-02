package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(password string) string  {
    // Create a new SHA-256 hash
    hasher := sha256.New()

    // Write the string to the hasher
    hasher.Write([]byte(password))

    // Compute the hash and get the byte slice
    hashBytes := hasher.Sum(nil)

    // Convert the byte slice to a hexadecimal string
    hashString := hex.EncodeToString(hashBytes)

	return hashString
}