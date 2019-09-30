package main

import (
	"encoding/hex"
	"os"
)

func getSecret() []byte {
	secret := os.Getenv("SECRET")
	if secret == "" {
		panic("Error: Must provide a secret key under env variable SECRET")
	}

	secretbite, err := hex.DecodeString(secret)

	if err != nil {
		panic(err)
	}

	return secretbite
}
