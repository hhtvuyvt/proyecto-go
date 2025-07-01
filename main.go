package main

import (
	"log"
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/routes"
)

func main() {
	r := routes.Router()

	log.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
