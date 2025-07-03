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

	sqlDB := db.Open()
	repo := models.BookRepository{DB: sqlDB}
	bookH := handlers.BookHandler{Repo: repo}

	// Login público
	mux.HandleFunc("/api/login", handlers.Login)

	// Protegido con middleware JWT
	mux.Handle("/api/books", middlewares.AuthMiddleware(http.HandlerFunc(bookH.Books)))
	mux.Handle("/api/books/", middlewares.AuthMiddleware(http.HandlerFunc(bookH.Book)))
	mux.Handle("/api/upload", middlewares.AuthMiddleware(http.HandlerFunc(handlers.UploadImage)))

	// Archivos estáticos
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusTemporaryRedirect))

	return middlewares.LoggerMiddleware(mux)
}
