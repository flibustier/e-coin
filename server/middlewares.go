package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/auth0-community/go-auth0"
	"github.com/gorilla/context"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

// getValidator returns the auth0 validator for the middleware
func getValidator() *auth0.JWTValidator {
	issuer := "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
	jwksURI := issuer + ".well-known/jwks.json"
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: jwksURI}, nil)

	aud := os.Getenv("AUTH0_AUDIENCE")
	audience := []string{aud}

	configuration := auth0.NewConfiguration(client, audience, issuer, jose.RS256)
	return auth0.NewValidator(configuration, nil)
}

// deny writes a unauthorized response
func deny(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode("Missing or invalid token.")
}

// authMiddleware assures that the access_token is valid
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := getValidator().ValidateRequest(r)

		if err != nil {
			log.Println("[ERROR] Token is not valid or missing token")
			deny(w)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// idMiddleware extracts id_token from request header and store the user email in context
func idMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Maybe we should request the id_token to auth0?
		idToken, err := jwt.ParseSigned(r.Header.Get("id_token"))
		if err != nil {
			log.Println("[ERROR] id_token cannot be parsed")
			deny(w)
			return
		}

		claims := map[string]interface{}{}
		err = getValidator().Claims(r, idToken, &claims)
		if err != nil {
			log.Println("[ERROR] No claim found in JWT")
			deny(w)
			return
		}

		email := claims["email"]
		if email != nil /*&& claims["email_verified"] != nil && claims["email_verified"].(bool)*/ {
			// Store email in context
			context.Set(r, "email", email.(string))
			next.ServeHTTP(w, r)
		} else {
			msg := "Email not found in JWT or email not verified"
			log.Println(msg)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(msg)
		}
	})
}
