package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/middlewares"
	"github.com/hhtvuyvt/proyecto-go/models"
)

// RouterConfig agrupa dependencias del router.
type RouterConfig struct {
	BookRepo models.BookRepository

	JWTKey []byte
}

// Router configura rutas HTTP.
func Router(
	cfg RouterConfig,
) http.Handler {

	mux := http.NewServeMux()

	bookHandler :=
		handlers.BookHandler{
			Repo: cfg.BookRepo,
		}

	// =====================
	// Publicas
	// =====================

	mux.HandleFunc(
		"/api/books",
		bookHandler.Books,
	)

	mux.HandleFunc(
		"/api/login",
		handlers.LoginHandler,
	)

	// =====================
	// PROTEGIDAS
	// =====================

	mux.Handle(
		"/api/books/",
		middlewares.AuthMiddleware(
			cfg.JWTKey,
			http.HandlerFunc(
				bookHandler.Book,
			),
		),
	)

	mux.Handle(
		"/api/upload",
		middlewares.AuthMiddleware(
			cfg.JWTKey,
			http.HandlerFunc(
				handlers.UploadImage,
			),
		),
	)

	// =====================
	// Archivos
	// =====================

	staticFiles :=
		http.FileServer(
			http.Dir("./static"),
		)

	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			staticFiles,
		),
	)

	// Imágenes subidas
	//
	// Permite acceder a:
	//
	// /uploads/imagen.jpg
	uploads :=
		http.FileServer(
			http.Dir("./uploads"),
		)

	mux.Handle(
		"/uploads/",
		http.StripPrefix(
			"/uploads/",
			uploads,
		),
	)

	mux.Handle(
		"/",
		http.RedirectHandler(
			"/static/index.html",
			http.StatusTemporaryRedirect,
		),
	)

	return middlewares.RecoverMiddleware(
		middlewares.LoggerMiddleware(
			mux,
		),
	)
}
