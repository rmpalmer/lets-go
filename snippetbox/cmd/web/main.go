package main

import (
	"fmt"
	"log"
	"net/http"
)

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
