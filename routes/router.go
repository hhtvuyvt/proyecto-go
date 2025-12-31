package routes

import (
	"net/http"
	"os"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/internal/db"
	"github.com/hhtvuyvt/proyecto-go/middlewares"
	"github.com/hhtvuyvt/proyecto-go/models"
)

// Router configura las rutas de la aplicación.
func Router() http.Handler {
	mux := http.NewServeMux()

	// Base de datos
	sqlDB := db.Open()
	repo := models.BookRepository{DB: sqlDB}
	bookHandler := handlers.BookHandler{Repo: repo}

	// JWT secret (inyectado)
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	// =========================
	// 🌐 RUTAS PÚBLICAS
	// =========================

	mux.HandleFunc("/api/books", bookHandler.Books)

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusTemporaryRedirect))

	// =========================
	// 🔐 RUTAS PROTEGIDAS
	// =========================

	mux.Handle(
		"/api/books/",
		middlewares.AuthMiddleware(jwtKey, http.HandlerFunc(bookHandler.Book)),
	)

	mux.Handle(
		"/api/upload",
		middlewares.AuthMiddleware(jwtKey, http.HandlerFunc(handlers.UploadImage)),
	)

	return mux
}
