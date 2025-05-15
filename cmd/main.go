package main

import (
	"log"
	"net/http"
)

type Application struct{}

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
