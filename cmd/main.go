package main

import (
	"log"
	"net/http"
)

type Application struct{}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	log.Printf("Server is running at PORT : %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
