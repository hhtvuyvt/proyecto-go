package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/middlewares"
	"github.com/hhtvuyvt/proyecto-go/models"
)

// RouterConfig agrupa las dependencias
// necesarias para construir el router.
type RouterConfig struct {
	BookRepo models.BookRepositoryInterface

	UserRepo models.UserRepositoryInterface

	JWTKey []byte
}

// Router configura todas las rutas HTTP.
func Router(
	cfg RouterConfig,
) http.Handler {

	mux :=
		http.NewServeMux()

	bookHandler :=
		handlers.BookHandler{
			Repo: cfg.BookRepo,
		}

	authHandler :=
		handlers.AuthHandler{

			UserRepo: cfg.UserRepo,

			JWTKey: cfg.JWTKey,
		}

	// =====================
	// Públicas
	// =====================

	mux.HandleFunc(
		"/api/books",
		bookHandler.Books,
	)

	mux.HandleFunc(
		"/api/login",
		authHandler.LoginHandler,
	)

	mux.HandleFunc(
		"/api/logout",
		handlers.LogoutHandler,
	)

	// =====================
	// Protegidas
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
	// Archivos estáticos
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

	// =====================
	// Imágenes subidas
	// =====================

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

	// =====================
	// Redirección raíz
	// =====================

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
