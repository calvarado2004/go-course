package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/calvarado2004/hello-template/pkg/config"
	"github.com/calvarado2004/hello-template/pkg/handlers"
	"github.com/calvarado2004/hello-template/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	//get the template cache from the app config
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/hello-world", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
