package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func Understand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Understanding Towing")
}
