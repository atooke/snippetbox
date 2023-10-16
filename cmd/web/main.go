package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	//define new command line arg, default value of :4000
	addr := flag.String("addr", ":4000", "HTTP network address")

	//parses command line flag into addr, call before using addr otherwise default of 4000 will always be used.
	flag.Parse()

	//setup custom logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// use the http.hewservemux to create a new router
	// then register home func as the handler for the / URL pattern
	mux := http.NewServeMux()

	//create a file server which serves files out of the ./ui/static directory
	//note - stripPrefix function is used to remove the /static prefix from the URL path when looking for files
	// note -  will santize path passed in to prevent directory traversal attacks
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//use the mux.Handle() to register the file server as teh handler for all
	//URL paths that start with "/static/".  For matching paths, we strip the "/static" prefix before the request reaches the file server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Use Info method to log the starting server message at info severity
	//along with the listen address as an atrtibute
	logger.Info("starting server", slog.String("addr", *addr))

	//user http.ListenAndServe() to start the web server.  host:port
	err := http.ListenAndServe(*addr, mux)
	// no structured logging eqv of Fatal so use error & then exit
	logger.Error(err.Error())
	os.Exit(1)

}
