package helpers

import (
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
)

func secretKey() []byte {
	return []byte("lemonilo_test_engineer")
}

func CreateJWT(id uint) (string, error) {
	secret := secretKey()

	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    strconv.FormatUint(uint64(id), 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := token.SignedString(secret)

	return jwt, err
}
