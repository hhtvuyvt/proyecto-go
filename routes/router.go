package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
	"github.com/hhtvuyvt/proyecto-go/middlewares"
	"github.com/hhtvuyvt/proyecto-go/models"
)

// RouterConfig agrupa las dependencias necesarias para construir el router.
type RouterConfig struct {
	ItemRepo   models.ItemRepositoryInterface // Antes BookRepo
	UserRepo   models.UserRepositoryInterface
	DetailRepo models.DetailRepositoryInterface // Antes ChapterRepo
	JWTKey     []byte
}

// Router configura todas las rutas HTTP de la aplicación.
func Router(cfg RouterConfig) http.Handler {
	mux := http.NewServeMux()

	// Instanciación de handlers con inyección de dependencias
	itemHandler := handlers.ItemHandler{Repo: cfg.ItemRepo} // Antes BookHandler
	authHandler := handlers.AuthHandler{
		UserRepo: cfg.UserRepo,
		JWTKey:   cfg.JWTKey,
	}
	detailHandler := handlers.DetailHandler{Repo: cfg.DetailRepo} // Antes ChapterHandler

	// =====================
	// Rutas Públicas
	// =====================
	mux.HandleFunc("GET /api/items", itemHandler.List) // Antes: bookHandler.Books
	mux.HandleFunc("POST /api/login", authHandler.LoginHandler)
	mux.HandleFunc("POST /api/logout", authHandler.LogoutHandler)
	// mux.HandleFunc("POST /api/register", authHandler.RegisterHandler) // <-- DESBLOQUEAR PRONTO

	// Ruta pública para leer detalles (ej: variantes de producto o info de habitación)
	mux.HandleFunc("GET /api/items/{item_id}/details/{id}", detailHandler.Get)
	mux.HandleFunc("GET /api/items/{item_id}/details", detailHandler.List)

	// =====================
	// Rutas Protegidas (Requieren JWT)
	// =====================
	mux.Handle("GET /api/me", middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(authHandler.MeHandler)))

	// CRUD de Items (Protegido)
	mux.Handle("POST /api/items", middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(itemHandler.Create)))
	mux.Handle("GET /api/items/{id}", middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(itemHandler.Get)))
	mux.Handle("PUT /api/items/{id}", middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(itemHandler.Update)))
	mux.Handle("DELETE /api/items/{id}", middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(itemHandler.Delete)))

	// CRUD de Detalles (Protegido)
	mux.Handle("POST /api/items/{item_id}/details", middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(detailHandler.Create)))

	// Utilidades
	mux.Handle("/api/upload", middlewares.AuthMiddleware(cfg.JWTKey, http.HandlerFunc(handlers.UploadImage)))

	// =====================
	// Archivos Estáticos e Imágenes
	// =====================
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// =====================
	// Redirección Raíz (Frontend)
	// =====================
	mux.Handle("/", http.RedirectHandler("/static/index.html", http.StatusTemporaryRedirect))

	// Aplicar Middlewares globales (Logger y Recover ante panics)
	return middlewares.RecoverMiddleware(
		middlewares.LoggerMiddleware(mux),
	)
}
