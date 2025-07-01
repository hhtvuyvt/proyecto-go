package main

import (
	"log"
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/routes"
)

func main() {
	log.Println("Servidor en http://localhost:8080")
	if err := http.ListenAndServe(":8080", routes.Router()); err != nil {
		log.Fatal(err)
	}
}
