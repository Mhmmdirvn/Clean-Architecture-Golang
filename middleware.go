package main

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"

	"Clean-Architecture/modules/login"
)

func MiddlewareJWTAuthorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenString, &login.MyClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("MySecret"), nil
		})

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		_, ok := token.Claims.(*login.MyClaims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token Invalid"))
			return
		}

		next(w, r)
	}
}