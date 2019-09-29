package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

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

}
