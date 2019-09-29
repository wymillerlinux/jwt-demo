package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my JWT Demo!")
	fmt.Fprintln(w, "Written by Wyatt J. Miller, 2019")
	fmt.Fprintln(w, "All rights reserved, licensed by the MPL 2.0")
	fmt.Fprintln(w, "Please go read the license")
}
