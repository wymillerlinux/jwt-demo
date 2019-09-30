package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

// in memory map to store passwords and username
// TODO: add database later
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
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
		fmt.Fprintln(w, "400 Bad Request. Please try again.")
		return
	}

	// get the password from in memory map
	// TODO: add database
	expectedPassword, ok := users[creds.Username]

	// if the password exists for the given user and if it's the same password we received,
	// move on my wayward son. if not, return an HTTP 401 (unauthorized)
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "401 Unauthorized. Please enter the correct credentials.")
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
		fmt.Fprintln(w, "500 Internal Server Error.")
		fmt.Fprintln(w, "We're getting to the bottom of the matter.")
		fmt.Fprintln(w, "Head over the status page while you wait.")
		return
	}

	// bake the cookie to send to the browser
	cookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}

	// set an HTTP cookie for the token
	http.SetCookie(w, cookie)
}
