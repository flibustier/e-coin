package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/auth0-community/go-auth0"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

/**
 * Get the auth0 validator
 */
func getValidator() *auth0.JWTValidator {
	issuer := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
	jwksURI := issuer + ".well-known/jwks.json"
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: jwksURI})

	aud := os.Getenv("AUTH0_AUDIENCE")
	audience := []string{aud}

	configuration := auth0.NewConfiguration(client, audience, issuer, jose.RS256)
	return auth0.NewValidator(configuration)
}

/**
 * This middleware assure that the access_token is valid
 */
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		_, err := getValidator().ValidateRequest(r)

		if err != nil {
			log.Println("[ERROR] Token is not valid or missing token")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Missing or invalid token.")

		} else {
			next.ServeHTTP(w, r)
		}
	})
}

/**
 * This function extract id_token from request headers and get the user email
 */
func getUserEmail(r *http.Request) (string, error) {
	// Maybe we should request the id_token to auth0?
	idToken, err := jwt.ParseSigned(r.Header.Get("id_token"))
	if err != nil {
		log.Println("[ERROR] id_token cannot be parsed")
		return "", err
	}

	claims := map[string]interface{}{}
	err = getValidator().Claims(r, idToken, &claims)
	if err != nil {
		log.Println("[ERROR] No claim found in JWT")
		return "", err
	}

	email := claims["email"]
	if email != nil && claims["email_verified"] != nil && claims["email_verified"].(bool) {
		return email.(string), nil
	}
	return "", fmt.Errorf("email not found in JWT or email not verified")
}
