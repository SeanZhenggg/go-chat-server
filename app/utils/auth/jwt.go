package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	SECRET_KEY []byte = []byte("jwt-token-secret")
)

type Claims struct {
	Account string
	jwt.RegisteredClaims
}

func TokenGenerate(account string) (string, error) {
	now := time.Now()

	claims := &Claims{
		Account: account,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(30 * time.Minute)),
			Audience:  jwt.ClaimStrings{account},
			ID:        account + strconv.FormatInt(now.Unix(), 10),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   account,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := tokenClaims.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return token, nil
}

func TokenValidation(token string) (string, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := tokenClaims.Claims.(*Claims)

	if !ok || claims.Account == "" {
		return "", err
	}

	return claims.Account, nil
}
