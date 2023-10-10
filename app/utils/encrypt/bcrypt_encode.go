package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordEncrypt(pwd string) (string, error) {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(encrypt), nil
}

func PasswordCompare(hash []byte, pwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pwd))

	if err != nil {
		return false, err
	}

	return true, nil
}
