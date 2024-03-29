package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kadekbagus/go-belajar/pkg/config"
	"github.com/kadekbagus/go-belajar/pkg/handlers"
	"github.com/kadekbagus/go-belajar/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
