package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	SecretKey = []byte("secret")
)

// GenerateTocken generates a jwt tocken and assign a username to it's claims and return it

func GenerateTokenString(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	var mapClaims jwt.MapClaims
	mapClaims = token.Claims.(jwt.MapClaims)
	mapClaims["username"] = username
	mapClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error on Generating key")
		return "", err
	}
	return tokenString, nil
}

// ParseToken parses a jwt token and returns the username in it's claims
func ParseTokenString(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
