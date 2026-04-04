package main

import (
	"log"
	"net/http"
	"html/template"
)

var templates *template.Template

type user struct {
	Nome string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func (w http.ResponseWriter, r *http.Request){
		u := user{
			"Daniel",
			"dandan@gmail.com",
		}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}