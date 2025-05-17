package main

import (
	"log"
	"net/http"

	"github.com/munnaMia/from-list/internals/model"
)

type Application struct {
	User model.User
}

func main() {
	app := &Application{}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: app.route(),
	}

	log.Printf("Server is running at PORT : %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
