package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("secret")

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString(jwtSecret)
}
