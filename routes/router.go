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

	mux.Handle("/api/books", http.HandlerFunc(bookH.Books))
	mux.Handle("/api/books/", http.HandlerFunc(bookH.Book))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusTemporaryRedirect))

	handler := middlewares.RecoverMiddleware(mux)
	handler = middlewares.LoggerMiddleware(handler)

	return handler
}
