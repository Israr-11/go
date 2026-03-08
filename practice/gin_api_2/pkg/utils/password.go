/*
WHAT THIS FILE DOES (Simple Explanation)

This file provides helper functions for password security.

Instead of storing passwords directly in the database, we:

1. Hash the password (convert it into a secure irreversible string)
2. Store the hashed password in the database
3. When user logs in, compare the entered password with the stored hash

Flow example:

User registers
   ↓
HashPassword("mypassword")
   ↓
Store hashed password in DB

User logs in
   ↓
CheckPasswordHash("mypassword", storedHash)
   ↓
Returns true or false

bcrypt is a secure hashing algorithm commonly used for passwords.
*/

package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes plain text password
func HashPassword(password string) (string, error) {

    // bcrypt.GenerateFromPassword requires password as []byte
    // []byte(password) converts the string into bytes
    // Think of this as "convert string format into raw data format"

    // 14 is the "cost factor"
    // Higher number = stronger but slower hashing
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

    // bytes is returned as []byte so we convert it back to string
    return string(bytes), err
}

// CheckPasswordHash compares password with hash
func CheckPasswordHash(password, hash string) bool {

    // bcrypt.CompareHashAndPassword checks if the password matches the stored 
	// hash, it returns nil if they match, otherwise returns an error

    err := bcrypt.CompareHashAndPassword(
        []byte(hash),     // stored hashed password from database
        []byte(password), // password user just entered
    )

    // If error is nil → password matched
    return err == nil
}