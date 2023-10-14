package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	//extract id parameter from query string & try to convert to int via atoi, if it can't be converted or is less than 1 return 404
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// use fmt.fprint to write the ID with response
	fmt.Fprint(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// check request method to ensure post method
	if r.Method != http.MethodPost {
		//if not post return 405 error on response
		// note: only possible to call w.WriteHeader() once per response, and after the status code has been written it can’t be changed
		//If you don’t call w.WriteHeader() explicitly, then the first call to w.Write() will automaticallysenda200 OKstatuscodetotheuser
		w.Header().Set("Allow", http.MethodPost) // add methods allowed in response
		http.Error(w, "Method Not Allowed", 405)
		return
	}
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
