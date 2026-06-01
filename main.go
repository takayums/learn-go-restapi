package main

import (
	"fmt"
	"net/http"
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
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/index", handleHello)

	address := "localhost:9000"
	fmt.Printf("server running on %s is yes\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
