package main

import (
	"fmt"
	"joshuamURD/testing/internal/config"
	"joshuamURD/testing/internal/handlers"
	"joshuamURD/testing/internal/renderer"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, _ := renderer.CreateTemplateCache()
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.IndexPage)

	fmt.Println("Listening on port", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}