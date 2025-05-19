package main

import "html/template"

func (app *Application) ParseTemp(paths []string) *template.Template {
	return template.Must(template.ParseFiles(paths...))
}