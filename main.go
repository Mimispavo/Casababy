package main

import (
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", home)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}
