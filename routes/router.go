package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/internal/db"
	"github.com/hhtvuyvt/proyecto-go/middlewares"
	"github.com/hhtvuyvt/proyecto-go/models"
)

// Router configura las rutas de la aplicación.
func Router() http.Handler {
	mux := http.NewServeMux()

	// Base de datos y repositorio
	sqlDB := db.Open()
	repo := models.BookRepository{DB: sqlDB}
	bookHandler := handlers.BookHandler{Repo: repo}

	// =========================
	// 🌐 RUTAS PÚBLICAS
	// =========================

	// Listar libros (GET)
	mux.HandleFunc("/api/books", bookHandler.Books)

	// Archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Página principal
	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusTemporaryRedirect))

	// =========================
	// 🔐 RUTAS PROTEGIDAS
	// =========================

	// Crear / Editar / Eliminar libro
	mux.Handle(
		"/api/books/",
		middlewares.AuthMiddleware(http.HandlerFunc(bookHandler.Book)),
	)

	// Subida de imágenes
	mux.Handle(
		"/api/upload",
		middlewares.AuthMiddleware(http.HandlerFunc(handlers.UploadImage)),
	)

	return mux
}
