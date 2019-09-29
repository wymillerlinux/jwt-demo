package main

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	// get the cookie from cookie jar
	cookie, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get the JWT string from the cookie
	// imagine getting your fortune out of a fortune cookie
	// also, initalize new instance of the claims struct
	tokenString := cookie.Value
	claims := &claims{}

	token, err := jwt.ParseWithClaims(tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	// if the parsed token came back with an error, either return
	// a HTTP 401 or a HTTP 400
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if the parsed token isn't valid, return a HTTP 401 (unauthorized)
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// welcome the user if the token is indeed valid
	w.Write([]byte(fmt.Sprintf("Welcome %s", claims.Username)))

}
