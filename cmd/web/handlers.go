package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	templateFiles := []string{"./ui/html/pages/base.tmpl.html", "./ui/html/pages/home.tmpl.html", "./ui/html/partials/nav.tmpl.html"}
	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error1", 500)
		return
	}
	if err := ts.ExecuteTemplate(w, "base", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error2", 500)
	}
	// fmt.Fprint(w, "Welcome to snippetbox")
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("we are going to print a sample snippet here"))
	fmt.Fprintf(w, "we are eventually going to print the snippet of id: %d here", id)
}

func snippetWrite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Error 405: Method not allowed "))
		http.Error(w, "The Request Method is  not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("we are going to create snippets with this route"))
}
