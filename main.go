package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/hhtvuyvt/proyecto-go/internal/db"
	"github.com/hhtvuyvt/proyecto-go/models"
	"github.com/hhtvuyvt/proyecto-go/routes"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró .env, usando variables del sistema")
	}

	// Configuración
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET no configurado")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH no configurado")
	}

	// Base de datos
	sqlDB := db.Open(dbPath)

	// Repositorios
	bookRepo := models.BookRepository{DB: sqlDB}

	// Router con dependencias inyectadas
	router := routes.Router(routes.RouterConfig{
		BookRepo: bookRepo,
		JWTKey:   []byte(jwtSecret),
	})

	log.Printf("Servidor en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
