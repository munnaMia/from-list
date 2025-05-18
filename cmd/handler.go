package main

import (
	"html/template"
	"net/http"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	// Page render data.
	data := map[string]any{
		"Title": "Home Page",
	}

	tmpl := template.Must(template.ParseFiles("./ui/html/layout/base.html", "./ui/html/partials/header.html", "./ui/html/partials/footer.html", "./ui/html/partials/nab.html", "./ui/html/pages/home.html"))

	if err := tmpl.Execute(w, data); err != nil {
		panic(err)
	}

}
func (app *Application) form(w http.ResponseWriter, r *http.Request)  {}
func (app *Application) list(w http.ResponseWriter, r *http.Request)  {}
func (app *Application) about(w http.ResponseWriter, r *http.Request) {}
