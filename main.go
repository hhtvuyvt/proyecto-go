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

	// Cargar .env si existe
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró .env, usando variables del sistema")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ==========================
	// Configuración JWT
	// ==========================

	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {

		if os.Getenv("E2E") == "true" {

			log.Println(
				"Modo E2E: usando JWT_SECRET temporal",
			)

			jwtSecret = "test-secret-key"

		} else {

			log.Fatal(
				"JWT_SECRET no configurado",
			)

		}
	}

	// ==========================
	// Base de datos
	// ==========================

	dbPath := os.Getenv("DB_PATH")

	if dbPath == "" {

		if os.Getenv("E2E") == "true" {

			log.Println(
				"Modo E2E: usando base de datos de prueba",
			)

			dbPath = "./test.db"

		} else {

			log.Fatal(
				"DB_PATH no configurado",
			)
		}
	}

	sqlDB := db.Open(dbPath)

	bookRepo := models.BookRepository{
		DB: sqlDB,
	}

	router := routes.Router(
		routes.RouterConfig{
			BookRepo: bookRepo,
			JWTKey:   []byte(jwtSecret),
		},
	)

	log.Printf(
		"servidor en http://localhost:%s",
		port,
	)

	log.Fatal(
		http.ListenAndServe(
			":"+port,
			router,
		),
	)
}
