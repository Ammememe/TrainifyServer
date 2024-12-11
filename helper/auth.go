package helper

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a given password using bcrypt.
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Error in hashing password")
	}
	return string(hash)
}

// CheckPasswordHash compares a plain password with its hashed version.
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

