package internal

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func Validate(str string) (isValid bool, token *jwt.Token, err error) {
	token, err = jwt.Parse(str, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing algorithm was used")
		}

		return []byte(os.Getenv("ADMIN_TOKEN_SECRET")), nil
	})

	if err != nil {
		return false, nil, err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, nil, fmt.Errorf("invalid token")
	}

	return true, token, nil
}
