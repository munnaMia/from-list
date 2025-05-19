package main

import (
	"net/http"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	files := []string{
		"./ui/html/layout/base.html",
		"./ui/html/partials/header.html",
		"./ui/html/partials/footer.html",
		"./ui/html/partials/nab.html",
		"./ui/html/pages/home.html",
	}

	// Page render data.
	data := map[string]any{
		"Title": "Home Page",
	}

	tmpl := app.ParseTemp(files)

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		panic(err)
	}

}
func (app *Application) form(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/layout/base.html",
		"./ui/html/partials/header.html",
		"./ui/html/partials/footer.html",
		"./ui/html/partials/nab.html",
		"./ui/html/pages/form.html",
	}

	// Page render data.
	data := map[string]any{
		"Title": "Form Page",
	}

	tmpl := app.ParseTemp(files)

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		panic(err)
	}
}
func (app *Application) list(w http.ResponseWriter, r *http.Request) {
	// Page render data.
	data := map[string]any{
		"Title": "List Page",
	}

	files := []string{
		"./ui/html/layout/base.html",
		"./ui/html/partials/header.html",
		"./ui/html/partials/footer.html",
		"./ui/html/partials/nab.html",
		"./ui/html/pages/list.html",
	}

	tmpl := app.ParseTemp(files)

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		panic(err)
	}
}
func (app *Application) about(w http.ResponseWriter, r *http.Request) {
	// Page render data.
	data := map[string]any{
		"Title": "About Page",
	}
	files := []string{
		"./ui/html/layout/base.html",
		"./ui/html/partials/header.html",
		"./ui/html/partials/footer.html",
		"./ui/html/partials/nab.html",
		"./ui/html/pages/about.html",
	}

	tmpl := app.ParseTemp(files)

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		panic(err)
	}
}
