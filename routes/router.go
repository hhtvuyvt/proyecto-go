package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/middlewares"
	"github.com/hhtvuyvt/proyecto-go/models"
)

// RouterConfig agrupa las dependencias necesarias para construir el router.
type RouterConfig struct {
	BookRepo models.BookRepositoryInterface
	UserRepo models.UserRepositoryInterface
	JWTKey   []byte
}

// Router configura todas las rutas HTTP de la aplicación.
func Router(cfg RouterConfig) http.Handler {
	mux := http.NewServeMux()

	bookHandler := handlers.BookHandler{Repo: cfg.BookRepo}
	authHandler := handlers.AuthHandler{
		UserRepo: cfg.UserRepo,
		JWTKey:   cfg.JWTKey,
	}

	// =====================
	// Rutas Públicas
	// =====================
	mux.HandleFunc("/api/books", bookHandler.Books)
	mux.HandleFunc("/api/login", authHandler.LoginHandler)
	mux.HandleFunc("/api/logout", authHandler.LogoutHandler)

	// =====================
	// Rutas Protegidas
	// =====================
	mux.Handle("/api/books/", middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(bookHandler.Book)))
	mux.Handle("/api/upload", middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(handlers.UploadImage)))

	// =====================
	// Archivos Estáticos e Imágenes
	// =====================
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// =====================
	// Redirección Raíz
	// =====================
	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusTemporaryRedirect))

	// Aplicar Middlewares globales
	return middlewares.RecoverMiddleware(
		middlewares.LoggerMiddleware(mux),
	)
}
