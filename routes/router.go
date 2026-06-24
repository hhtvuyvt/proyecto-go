package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/middlewares"
	"github.com/hhtvuyvt/proyecto-go/models"
)

// RouterConfig agrupa todas las dependencias necesarias
// para construir el router.
//
// BookRepo:
// repositorio encargado del acceso a libros.
//
// JWTKey:
// clave usada para firmar y validar tokens JWT.
type RouterConfig struct {
	BookRepo models.BookRepository
	JWTKey   []byte
}

// Router configura todas las rutas HTTP de la aplicación.
func Router(cfg RouterConfig) http.Handler {

	mux := http.NewServeMux()

	bookHandler :=
		handlers.BookHandler{
			Repo: cfg.BookRepo,
		}

	// ==========================
	// RUTAS PÚBLICAS
	// ==========================

	// Obtener y crear libros.
	//
	// GET  /api/books
	// POST /api/books
	mux.HandleFunc(
		"/api/books",
		bookHandler.Books,
	)

	// Login público.
	//
	// Genera un JWT que el frontend
	// usará para operaciones protegidas.
	mux.HandleFunc(
		"/api/login",
		handlers.LoginHandler,
	)

	// ==========================
	// RUTAS PROTEGIDAS
	// ==========================

	// Editar y borrar libros.
	//
	// PUT    /api/books/{id}
	// DELETE /api/books/{id}
	//
	// Estas rutas requieren:
	//
	// Authorization:
	// Bearer <token>
	mux.Handle(
		"/api/books/",
		middlewares.AuthMiddleware(
			cfg.JWTKey,
			http.HandlerFunc(
				bookHandler.Book,
			),
		),
	)

	// Subida de imágenes protegida.
	mux.Handle(
		"/api/upload",
		middlewares.AuthMiddleware(
			cfg.JWTKey,
			http.HandlerFunc(
				handlers.UploadImage,
			),
		),
	)

	// ==========================
	// ARCHIVOS ESTÁTICOS
	// ==========================

	fs :=
		http.FileServer(
			http.Dir("./static"),
		)

	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			fs,
		),
	)

	// Entrada principal.
	mux.Handle(
		"/",
		http.RedirectHandler(
			"/static/index.html",
			http.StatusTemporaryRedirect,
		),
	)

	return mux
}
