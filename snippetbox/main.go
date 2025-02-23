package main

import (
	"fmt"
	"log"
	"net/http"
)

// define a home handler function which writes a byte slice
// containing a message as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("diplay a form for creating a snippet"))
}

func main() {
	fmt.Println("hello from my application")

	// use the newservemux function to initialize a servemux
	// and register the home function as the handler for the "/" pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	//print a log message about starting
	log.Print("starting serve on :4000")

	//start a new wev server, using the TCP network address and the servemux
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

	fmt.Println("Goodbye from my application")
}
