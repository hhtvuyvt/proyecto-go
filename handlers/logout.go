package handlers

import (
	"encoding/json"
	"net/http"
)

// LogoutHandler elimina la cookie de autenticación.
//
// El JWT no se invalida en el servidor porque actualmente
// es un token stateless. Simplemente se elimina la cookie
// del navegador para finalizar la sesión.
func LogoutHandler(
	w http.ResponseWriter,
	_ *http.Request,
) {

	http.SetCookie(
		w,
		&http.Cookie{

			Name: "token",

			Value: "",

			Path: "/",

			HttpOnly: true,

			SameSite: http.SameSiteLaxMode,

			// Debe coincidir con LoginHandler.
			Secure: false,

			// Expira inmediatamente.
			MaxAge: -1,
		},
	)

	if err :=
		json.NewEncoder(w).Encode(
			map[string]string{
				"message": "logout correcto",
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
