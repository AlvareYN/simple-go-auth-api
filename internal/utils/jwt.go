package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		return claims, nil
	}

	log.Println(err)
	return nil, err
}

func GenerateToken(claims map[string]interface{}) (string, error) {
	var jwtClaims jwt.Claims = jwt.MapClaims{}

	for key, value := range claims {
		jwtClaims.(jwt.MapClaims)[key] = value
	}

	jwtClaims.(jwt.MapClaims)["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	tokenString, err := jwtToken.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
