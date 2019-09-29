package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func refresh(w http.ResponseWriter, r *http.Request) {
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

	// elasped time check, return HTTP 400 if the token is too fresh
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create brand new token
	expirationTime := time.Now().Add(10 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := newToken.SignedString(jwtKey)

	// if the new token comes back with an error, return a HTTP 500
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// bake a brand new cookie
	NewCookie := &http.Cookie{
		Name:    "token",
		Value:   newTokenString,
		Expires: expirationTime,
	}

	// set the brand new cookie on the cooling rack
	http.SetCookie(w, NewCookie)
}
