package controllers

import (
	//	"fmt"
	"../../gophers/plate"
	_ "log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	server := plate.NewServer()

	tmpl, _ := server.Template(w)

	tmpl.Template = "templates/home/index.html"

	tmpl.DisplayTemplate()
}
