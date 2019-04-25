package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"autoapp/web/app"
)

const defaultPort = "8080"
const defaultName = "Gopher"

func main() {
	name, exists := os.LookupEnv("NAME")
	if !exists {
		name = defaultName
	}

	s := app.Server{
		Router: http.NewServeMux(),
		Welcome: &app.Welcome{
			Name: name,
			Time: time.Now().Format(time.Stamp),
		},
	}

	s.Routes("web")

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = defaultPort
	}

	log.Println("Server listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, s.Router))
}
