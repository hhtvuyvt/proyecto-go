package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/internal/db"
	"github.com/hhtvuyvt/proyecto-go/middlewares"
	"github.com/hhtvuyvt/proyecto-go/models"
)

// Router configura las rutas de la aplicación web y retorna el enrutador principal.
func Router() http.Handler {
	mux := http.NewServeMux()

	// Conexión a la base de datos
	sqlDB := db.Open()
	repo := models.BookRepository{DB: sqlDB}
	bookHandler := handlers.BookHandler{Repo: repo}

	// Rutas API
	mux.HandleFunc("/api/books", bookHandler.Books)
	mux.HandleFunc("/api/books/", bookHandler.Book)
	mux.HandleFunc("/api/upload", handlers.UploadImage)

	// Archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Redirección a la página principal
	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusTemporaryRedirect))

	// Middleware de autenticación aplicado al router
	return middlewares.AuthMiddleware(mux)
}
