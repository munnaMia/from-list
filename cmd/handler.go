package main

import (
	"fmt"
	"net/http"

	"github.com/munnaMia/from-list/internals/utils"
)

const DBpath = "internals/storage/formDb.json"

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

func (app *Application) fromInput(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allow", http.StatusMethodNotAllowed)
		return
	}

	// store that into a JSON db
	dataModel := app.User

	dataModel.Name = r.PostFormValue("user_name")
	dataModel.Email = r.PostFormValue("user_email")
	dataModel.Gender = r.PostFormValue("user_gender")
	dataModel.Password = r.PostFormValue("user_password")
	fmt.Println(dataModel) // temporary

	err := utils.CreateIfNotExist(DBpath)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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
