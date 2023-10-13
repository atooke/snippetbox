package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// responseWriter - parameter provides methods for assembling a HTTP response and sending it to the user
	// request - parameter is a pointer to a struct which holds information about the current request
	if r.URL.Path != "/" {
		http.NotFound(w, r) //returns 404 error
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

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
