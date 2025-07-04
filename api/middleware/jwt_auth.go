package middleware

import (
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

var (
	TokenAuth = jwtauth.New("HS256", []byte("your-secret-key"), nil)
)

func JWTMiddleware(auth *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return jwtauth.Verifier(auth)(jwtauth.Authenticator(auth)(next))
	}
}
