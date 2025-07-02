package middlewares

import (
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware imprime la URL y el tiempo de respuesta
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("⏳ %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("✅ %s %s (%s)", r.Method, r.URL.Path, time.Since(start))
	})
}

// RecoverMiddleware evita que el servidor se caiga por panics
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("💥 Panic: %v", err)
				http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
