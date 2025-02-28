package logic

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) ([]byte, error) {
	for i := range password {
		password[i] ^= 0xAA
	}

	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	for i := range password {
		password[i] = 0
	}

	return hashed, nil
}

func ComparePasswords(hashedPassword, password []byte) error {
	for i := range password {
		password[i] ^= 0xAA
	}

	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err
	}

	for i := range password {
		password[i] = 0
	}

	return nil
}
