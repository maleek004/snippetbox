package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to snippetbox")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	port := ":4000"
	log.Printf("Server starting at port %s", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
