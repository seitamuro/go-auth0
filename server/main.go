// server/main.go
// APIのエントリポイント

package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/seitamuro/go-auth0/server/handlers/v1"
)

const (
	port = 8000
)

func main() {
	mux := http.NewServeMux()
	// /v1へのリクエストが来た場合のハンドラを追加
	mux.HandleFunc("/v1", v1.HandleIndex)

	addr := fmt.Sprintf(":%d", port)

	// localhost:8000 でサーバーを立ち上げる
	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
