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
