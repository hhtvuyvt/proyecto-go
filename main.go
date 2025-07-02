package main

import (
	"log"
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/routes"
)

func main() {
	log.Println("Servidor en http://localhost:8081")
	if err := http.ListenAndServe(":8081", routes.Router()); err != nil {
		log.Fatal(err)
	}
}
