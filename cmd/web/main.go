package main

import (
	"log"
	"net/http"
)

func main() {
	// use the http.hewservemux to create a new router
	// then register home func as the handler for the / URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// pri9nt a log message to say server is starting
	log.Print("starting server on :4000")

	//user http.ListenAndServe() to start the web server.  host:port
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
