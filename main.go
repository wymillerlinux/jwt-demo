package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Setting up JWT microservice!")

	router := mux.NewRouter()
	router.HandleFunc("/signin", Signin).Methods("POST")
	router.HandleFunc("/welcome", Welcome).Methods()
	router.HandleFunc("/refresh", Refresh).Methods()

	log.Println("Starting JWT microservice!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
