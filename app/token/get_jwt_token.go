package token

import (
	"net/http"
)

func GetToken(r *http.Request) bool {
	_, err := TokenValid(r)
	if err == nil {
		return true
	}
	return false
}
