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
	
	tmpl := template.Must(template.ParseFiles("./ui/html/home.html"))

	if err := tmpl.Execute(w, nil); err != nil {
		panic(err)
	}

}
func (app *Application) form(w http.ResponseWriter, r *http.Request)  {}
func (app *Application) list(w http.ResponseWriter, r *http.Request)  {}
func (app *Application) about(w http.ResponseWriter, r *http.Request) {}
