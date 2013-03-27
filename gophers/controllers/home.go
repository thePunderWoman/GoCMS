package controllers

import (
	//	"fmt"
	"../../gophers/plate"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	server := plate.NewServer()
	log.Println("This is a test")
	tmpl, _ := server.Template(w)

	tmpl.Template = "templates/home/index.html"

	tmpl.DisplayTemplate()
}
