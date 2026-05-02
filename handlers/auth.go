package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware valida el token JWT en las peticiones protegidas.
func AuthMiddleware(next http.Handler) http.Handler {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtKey) == 0 {
		log.Fatal("Falta JWT_SECRET en .env")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "token requerido", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "token inválido", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// LoginHandler genera un token simple (demo)
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "admin",
	})

	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "error generando token", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]string{"token": tokenStr}); err != nil {
		http.Error(w, "error al generar respuesta", http.StatusInternalServerError)
		return
	}
}