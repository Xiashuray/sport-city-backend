package token

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func TokenValid(r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("Authorization")
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil
	})
	return token, err
}
