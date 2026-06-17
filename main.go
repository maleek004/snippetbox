package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Welcome to snippetbox")
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("we are going to print a sample snippet here"))
}

func snippetWrite(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("we are going to create snippets with this route"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/write", snippetWrite)

	port := ":4000"
	log.Printf("Server starting at port %s", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
