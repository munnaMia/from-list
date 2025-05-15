package main

import "net/http"

func (app *Application) route() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/form", app.form)
	mux.HandleFunc("/list", app.list)
	mux.HandleFunc("/about", app.about)

	return mux
}
