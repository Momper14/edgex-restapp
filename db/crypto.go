package db

import "golang.org/x/crypto/bcrypt"

// hashAndSalt has and salts the password
func hashAndSalt(password *string) (*string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	str := string(hash)
	return &str, err
}
