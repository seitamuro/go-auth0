package auth0

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type JWKS struct {
	Keys []JSONWebKeys `json:"keys"`
}

func FetchJWKS(auth0Domain string) (*JWKS, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/.well-known/jwks.json", auth0Domain))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jwks := &JWKS{}
	err = json.NetDecoder(resp.Body).Decode(jwks)

	return jwks, err
}