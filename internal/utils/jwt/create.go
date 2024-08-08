package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(u string) (tokenString string, err error) {
	privatePEM, err := os.ReadFile("./cert/ecdsa-p521-private.pem")
	if err != nil {
		return "", err
	}

	secret, err := jwt.ParseECPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return "", err
	}

	claims := jwt.RegisteredClaims{
		Issuer:    "Agape",
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	return token.SignedString(secret)
}
