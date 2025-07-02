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

	// DB y repositorio
	sqlDB := db.Open()
	repo := models.BookRepository{DB: sqlDB}
	bookH := handlers.BookHandler{Repo: repo}

	// API
	mux.Handle("/api/books", http.HandlerFunc(bookH.Books))
	mux.Handle("/api/books/", http.HandlerFunc(bookH.Book))
	mux.HandleFunc("/api/upload", handlers.UploadImage) // 👈 NUEVO endpoint

	// Archivos estáticos
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))) // 👈 para servir imágenes

	// Página principal
	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusTemporaryRedirect))

	// Middleware (logger + recover)
	handler := middlewares.RecoverMiddleware(mux)
	handler = middlewares.LoggerMiddleware(handler)

	return handler
}
