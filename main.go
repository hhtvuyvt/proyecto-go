package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/hhtvuyvt/proyecto-go/internal/db"
	"github.com/hhtvuyvt/proyecto-go/models"
	"github.com/hhtvuyvt/proyecto-go/routes"
)

func main() {

	// ==========================
	// Variables de entorno
	// ==========================

	if err :=
		godotenv.Load(); err != nil {

		log.Println(
			"No se encontró .env, usando variables del sistema",
		)

	}

	// ==========================
	// Puerto
	// ==========================

	port :=
		os.Getenv(
			"PORT",
		)

	if port == "" {

		port = "8080"

	}

	// ==========================
	// JWT
	// ==========================

	jwtSecret :=
		os.Getenv(
			"JWT_SECRET",
		)

	if jwtSecret == "" {

		if os.Getenv("E2E") == "true" {

			log.Println(
				"Modo E2E: usando JWT_SECRET temporal",
			)

			jwtSecret =
				"test-secret-key"

		} else {

			log.Fatal(
				"JWT_SECRET no configurado",
			)

		}

	}

	// ==========================
	// Base de datos
	// ==========================

	dbPath :=
		os.Getenv(
			"DB_PATH",
		)

	if dbPath == "" {

		if os.Getenv("E2E") == "true" {

			log.Println(
				"Modo E2E: usando base de datos temporal",
			)

			dbPath =
				"./test.db"

		} else {

			log.Fatal(
				"DB_PATH no configurado",
			)

		}

	}

	sqlDB, err :=
		db.Open(
			dbPath,
		)

	if err != nil {

		log.Fatal(err)

	}

	defer func() {

		if err :=
			sqlDB.Close(); err != nil {

			log.Println(
				"Error cerrando base de datos:",
				err,
			)

		}

	}()

	// ==========================
	// Repositorios
	// ==========================

	bookRepo := &models.BookRepository{
		DB: sqlDB,
	}

	userRepo := &models.UserRepository{
		DB: sqlDB,
	}

	// ==========================
	// Router
	// ==========================

	router :=
		routes.Router(
			routes.RouterConfig{

				BookRepo: bookRepo,

				UserRepo: userRepo,

				JWTKey: []byte(
					jwtSecret,
				),
			},
		)

	log.Println("servidor iniciado en puerto " + port)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Fatal(
		server.ListenAndServe(),
	)
}
