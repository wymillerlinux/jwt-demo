package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

// in memory map to store passwords and username
// TODO: add database later
var users = map[string]string{
	"user": "password",
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func signin(w http.ResponseWriter, r *http.Request) {
	var creds credentials

	// decode JSON into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)

	// if there's an error with the JSON, return an HTTP 400 (bad request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get the password from in memory map
	// TODO: add database
	expectedPassword, ok := users[creds.Username]

	// if the password exists for the given user and if it's the same password we received,
	// move on my wayward son. if not, return an HTTP 401 (unauthorized)
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// set the expiration time on the token given to the user
	expirationTime := time.Now().Add(10 * time.Minute)

	// set up claims, includes username and expiry time
	claims := &claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// set up the actual token. if there's an error, there's something wrong
	// the actual generation so return an HTTP 500 (internal)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// set an HTTP cookie for the token
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
