package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// define a home handler function which writes a byte slice
// containing a message as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// extract the value of the id wildcard from the reqeust
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// interpolate the id value and give back as a message
	fmt.Fprintf(w, "Display a specific snippent with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("diplay a form for creating a new snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
