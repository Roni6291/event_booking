package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(name string, userId int64, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name":   name,
			"userId": userId,
			"exp":    time.Now().Add(time.Hour * 2).Unix(),
		},
	)
	return token.SignedString([]byte(key))
}

func VerifyToken(token, key string) error {
	parsedToken, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(key), nil
		},
	)
	if err != nil {
		return errors.New("couldn't parse token")

	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return errors.New("invalid token")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return errors.New("invalid token claims")
	// }
	// name, _ := claims["name"].(string)
	// userId, _ := claims["userId"].(int64)
	return nil
}
