package main

import "net/http"

func (app *Application) route() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/form", app.form)
	mux.HandleFunc("/forminput", app.fromInput)
	mux.HandleFunc("/list", app.list)
	mux.HandleFunc("/about", app.about)

	return mux
}
