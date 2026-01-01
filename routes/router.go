package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/middlewares"
	"github.com/hhtvuyvt/proyecto-go/models"
)

// RouterConfig agrupa todas las dependencias del router.
type RouterConfig struct {
	BookRepo models.BookRepository
	JWTKey   []byte
}

// Router configura las rutas de la aplicación.
func Router(cfg RouterConfig) http.Handler {
	mux := http.NewServeMux()

	bookHandler := handlers.BookHandler{Repo: cfg.BookRepo}

	// 🌐 Rutas públicas
	mux.HandleFunc("/api/books", bookHandler.Books)

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusTemporaryRedirect))

	// 🔐 Rutas protegidas
	mux.Handle(
		"/api/books/",
		middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(bookHandler.Book)),
	)

	mux.Handle(
		"/api/upload",
		middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(handlers.UploadImage)),
	)

	return mux
}
