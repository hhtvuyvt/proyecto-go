package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/internal/db"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	// Conexión BD y handler
	sqlDB := db.Open()
	bookH := handlers.BookHandler{DB: sqlDB}

	// API REST
	mux.HandleFunc("/api/books", bookH.Books)
	mux.HandleFunc("/api/books/", bookH.Book)

	// Archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Página principal
	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusTemporaryRedirect))
	return mux
}
