package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// LoginHandler genera un token JWT.
//
// Este handler solamente crea el token.
// La validación del token pertenece al paquete middlewares.
func LoginHandler(
	w http.ResponseWriter,
	_ *http.Request,
) {

	jwtKey := []byte(
		os.Getenv("JWT_SECRET"),
	)

	token :=
		jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			jwt.MapClaims{
				"user": "admin",
			},
		)

	tokenStr, err :=
		token.SignedString(jwtKey)

	if err != nil {

		http.Error(
			w,
			"error generando token",
			http.StatusInternalServerError,
		)

		return
	}

	if err :=
		json.NewEncoder(w).Encode(
			map[string]string{
				"token": tokenStr,
			},
		); err != nil {

		http.Error(
			w,
			"error generando respuesta",
			http.StatusInternalServerError,
		)

		return
	}
}
