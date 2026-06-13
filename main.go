package main

import (
	"html/template"
	"log"
	"net/http"
)

type M map[string]any

func main() {
	// Routing
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := M{"name": "Asraf"}
		tmpl := template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))

		err := tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := M{"name": "Asraf"}
		tmpl := template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))
		err := tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	address := "localhost:9000"
	log.Println("server running on localhost:9000")
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
