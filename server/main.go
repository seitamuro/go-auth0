// server/main.go
// APIのエントリポイント

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	v1 "github.com/seitamuro/go-auth0/server/handlers/v1"
	"github.com/seitamuro/go-auth0/server/handlers/v1/users/me"
	"github.com/seitamuro/go-auth0/server/middlewares/auth0"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("API_URL")
	domain := os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("AUTH0_CLIENT_ID")

	jwks, err := auth0.FetchJWKS(domain)
	if err != nil {
		log.Fatal(err)
	}

	jwtMiddleware, err := auth0.NewMiddleware(domain, clientID, jwks)
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/v1", v1.HandleIndex)
	mux.Handle("/v1/users/me", auth0.UseJWT(http.HandlerFunc(me.HandleIndex)))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	})

	wrappedMux := auth0.WithJWTMiddleware(jwtMiddleware)(mux)
	wrappedMux = c.Handler(wrappedMux)

	log.Printf("Listing on %s", url)
	if err := http.ListenAndServe(url, wrappedMux); err != nil {
		log.Fatal(err)
	}
}
