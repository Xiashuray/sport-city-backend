package token

import "net/http"

func Update_token(r *http.Request, uid string) (string, error) {
	_, err := TokenValid(r)
	if err == nil {
		return "token not expired", err
	}
	token, erro := CreateToken(uid)
	if erro == nil {
		return "error create", erro
	}
	return token, erro
}
