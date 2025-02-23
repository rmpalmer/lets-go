package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// define a home handler function which writes a byte slice
// containing a message as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// extract the value of the id wildcard from the reqeust
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// interpolate the id value and give back as a message
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("diplay a form for creating a snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new snippet..."))
}

func main() {
	fmt.Println("hello from my application")

	// use the newservemux function to initialize a servemux
	// and register the home function as the handler for the "/" pattern
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	//print a log message about starting
	log.Print("starting serve on :4000")

	//start a new wev server, using the TCP network address and the servemux
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

	fmt.Println("Goodbye from my application")
}
