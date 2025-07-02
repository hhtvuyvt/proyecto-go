package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/internal/db"
	"github.com/hhtvuyvt/proyecto-go/middlewares"
	"github.com/hhtvuyvt/proyecto-go/models"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	// Inicializar repositorio
	sqlDB := db.Open()
	repo := models.BookRepository{DB: sqlDB}
	bookH := handlers.BookHandler{Repo: repo}

	// API REST
	mux.Handle("/api/books", http.HandlerFunc(bookH.Books))
	mux.Handle("/api/books/", http.HandlerFunc(bookH.Book))

	// Archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Página principal
	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusTemporaryRedirect))

	// Aplicar middlewares
	handler := middlewares.RecoverMiddleware(mux)
	handler = middlewares.LoggerMiddleware(handler)

	return handler
}
