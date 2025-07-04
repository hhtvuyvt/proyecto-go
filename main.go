package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hhtvuyvt/proyecto-go/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Advertencia: no se pudo cargar .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor en http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, routes.Router()); err != nil {
		log.Fatal(err)
	}
}
