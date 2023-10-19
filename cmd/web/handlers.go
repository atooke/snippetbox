/*
The pattern that we’re using to inject dependencies won’t work if your handlers are spread across multiple packages. In that case,
an alternative approach is to create a standalone config package which exports an Application struct, and have your handler functions
close over this to form a closure. See -> https://gist.github.com/alexedwards/5cd712192b4831058b21
*/
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//read template file and if error return 500 error to user.
	//note - file path that you pass to the template.ParseFiles() function must either be relative to your current working directory, or an absolute path
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	//use Execute() on template to set to write content as response body, last parameter is dynamic data to pass to template.  nil for now
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Use the notFound() helper.
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper. return
	}
	w.Write([]byte("Create a new snippet..."))
}
