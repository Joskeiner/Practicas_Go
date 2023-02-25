package main

import (
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "COME HOME ")
}
