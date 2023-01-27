// ここでClaimsのペイロードのsubから認証済みのユーザーデータを取得する
// 本番環境ではここでDBなどにクエリを投げて識別子に対応するユーザーレコードを取得する
package me

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/form3tech-oss/jwt-go"
	"github.com/seitamuro/go-auth0/server/middlewares/auth0"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	subToUsers = map[string]User{
		"auth0|61a8178b21127500715968e2": {
			Name: "seitamuro",
			Age:  20,
		},
	}
)

func getUser(sub string) *User {
	user, ok := subToUsers[sub]
	if !ok {
		return nil
	}
	return &user
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	token := auth0.GetJWT(r.Context())
	fmt.Printf("jwt %+v\n", token)

	claims := token.Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)

	user := getUser(sub)
	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(res)
}
