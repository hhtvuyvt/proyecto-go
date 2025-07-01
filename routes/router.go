package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/internal/db"
	"github.com/hhtvuyvt/proyecto-go/internal/repository"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	// BD y repositorio
	sqlDB := db.Open()
	repo := repository.BookRepository{DB: sqlDB}
	bookH := handlers.BookHandler{Repo: repo}

	// API
	mux.HandleFunc("/api/books", bookH.Books)
	mux.HandleFunc("/api/books/", bookH.Book)

	// Archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Página principal
	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusTemporaryRedirect))
	return mux
}
