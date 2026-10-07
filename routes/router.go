package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/middlewares"
	"github.com/hhtvuyvt/proyecto-go/models"
)

// RouterConfig agrupa las dependencias necesarias para construir el router.
type RouterConfig struct {
	BookRepo    models.BookRepositoryInterface
	UserRepo    models.UserRepositoryInterface
	ChapterRepo models.ChapterRepositoryInterface
	JWTKey      []byte
}

// Router configura todas las rutas HTTP de la aplicación.
func Router(cfg RouterConfig) http.Handler {
	mux := http.NewServeMux()

	bookHandler := handlers.BookHandler{Repo: cfg.BookRepo}
	authHandler := handlers.AuthHandler{
		UserRepo: cfg.UserRepo,
		JWTKey:   cfg.JWTKey,
	}
	chapterHandler := handlers.ChapterHandler{ChapterRepo: cfg.ChapterRepo} // <-- Instanciado con el repo inyectado

	// =====================
	// Rutas Públicas
	// =====================
	mux.HandleFunc("GET /api/books", bookHandler.Books)
	mux.HandleFunc("POST /api/login", authHandler.LoginHandler)
	mux.HandleFunc("POST /api/logout", authHandler.LogoutHandler)

	// Ruta pública para leer capítulos
	mux.HandleFunc("GET /api/books/{book_id}/chapters/{id}", chapterHandler.GetChapterHandler)
	mux.HandleFunc("GET /api/books/{book_id}/chapters", chapterHandler.ListChaptersHandler)

	// =====================
	// Rutas Protegidas
	// =====================
	mux.Handle("GET /api/me", middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(authHandler.MeHandler)))
	mux.Handle("POST /api/books", middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(bookHandler.Books)))
	mux.Handle("POST /api/books/{book_id}/chapters", middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(chapterHandler.CreateChapterHandler)))
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
