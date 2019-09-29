package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my JWT Demo!")
	fmt.Fprintf(w, "Written by Wyatt J. Miller, 2019")
	fmt.Fprintf(w, "All rights reserved, licensed by the MPL 2.0")
	fmt.Fprintf(w, "Please go read the license")
}
