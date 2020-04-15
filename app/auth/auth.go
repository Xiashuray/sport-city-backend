package auth

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Auth_google(uid string) (user string) {
	opt := option.WithCredentialsFile("./config/googleToken.json")
	App, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		println(err)
	}
	client, err := App.Auth(context.Background())
	u, err := client.GetUser(context.Background(), uid)
	if err != nil {
		return "error user"
	}
	return u.UID
}

func Auth_server() {
	fmt.Println("dasd")
}
