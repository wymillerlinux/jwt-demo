package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Setting up JWT microservice!")

	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)

	log.Println("Starting JWT microservice!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
