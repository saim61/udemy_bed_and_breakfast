package main

import (
	"fmt"
	"log"
	"net/http"
	"udemy_bed_and_breakfast/pkg/config"
	"udemy_bed_and_breakfast/pkg/handlers"
	"udemy_bed_and_breakfast/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache ---> ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Staring application on port %s", portNumber)
	fmt.Print("\n")
	_ = http.ListenAndServe(portNumber, nil)
}
