package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	message := "Welcome"
	w.Write([]byte(message))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello World"
	w.Write([]byte(message))
}

func main() {
	// Static Assets
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	// Routing
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/index", handleHello)

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// Rendering HTML
		filepath := path.Join("views", "index.html")
		tmpl, err := template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"title": "Learning Golang Web",
			"name":  "Asraf",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	address := "localhost:9000"
	fmt.Printf("server running on %s is yes\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
