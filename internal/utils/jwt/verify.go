package jwt

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(str string) error {
	publicPEM, err := os.ReadFile("./cert/ecdsa-p521-public.pem")
	if err != nil {
		return err
	}

	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		key, err := jwt.ParseECPublicKeyFromPEM(publicPEM)
		if err != nil {
			return nil, err
		}

		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, err
		}

		return key, nil
	})
	if err != nil {
		return err
	}

	iss, err := token.Claims.GetIssuer()
	if err != nil {
		return err
	}

	if iss != "Agape" {
		return err
	}

	if !token.Valid {
		return err
	}

	return nil
}
