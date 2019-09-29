package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Setting up JWT microservice!")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", root)
	router.HandleFunc("/signin", signin)
	router.HandleFunc("/welcome", welcome)
	//router.HandleFunc("/refresh", refresh)

	log.Println("Starting JWT microservice!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
