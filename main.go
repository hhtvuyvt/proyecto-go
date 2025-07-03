package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/routes"
)

func main() {
	port := 8080
	log.Println("Servidor en http://localhost:{}", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), routes.Router()); err != nil {
		log.Fatal(err)
	}
}
