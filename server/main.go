package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	v1 "github.com/seitamuro/go-auth0/server/handlers/v1"
	"github.com/seitamuro/go-auth0/server/handlers/v1/users/me"
	"github.com/seitamuro/go-auth0/server/middlewares/auth0"
)

const (
	port     = 8000
	domain   = "<AUTH0_DOMAIN>"
	clientID = "<AUTH0_CLIENT_ID>"
)

func main() {
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

	addr := fmt.Sprintf(":%d", port)

	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		log.Fatal(err)
	}
}
