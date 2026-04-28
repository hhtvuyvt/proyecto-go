// auth.go define el endpoint de autenticación de usuarios.
// Retorna un token JWT para acceso autorizado a rutas protegidas.
package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// Credentials representa los datos enviados por el cliente en el login.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims son los datos que se almacenan en el token JWT.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// LoginHandler maneja POST /api/login y responde con un JWT si las credenciales son válidas.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "credenciales inválidas", http.StatusBadRequest)
		return
	}

	// Autenticación simplificada (puede ser reemplazada por consulta a BD).
	if creds.Username != "admin" || creds.Password != "1234" {
		http.Error(w, "usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	expiration := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "error generando token", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]string{"token": tokenStr}); err != nil {
        log.Printf("Error codificando JSON: %v", err) // ¡Esto te salva la vida en producción!
        http.Error(w, "error al generar respuesta", http.StatusInternalServerError)
        return
    }
}
