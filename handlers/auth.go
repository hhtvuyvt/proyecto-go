package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/hhtvuyvt/proyecto-go/models"
)

// AuthHandler maneja la autenticación.
type AuthHandler struct {
	UserRepo models.UserRepositoryInterface
	JWTKey   []byte
}

// LoginRequest representa el cuerpo JSON recibido
// desde el formulario de login.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler autentica al usuario
// y crea una cookie HttpOnly con el JWT.
func (h AuthHandler) LoginHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {

		http.Error(
			w,
			"método no permitido",
			http.StatusMethodNotAllowed,
		)

		return
	}

	var req LoginRequest

	if err :=
		json.NewDecoder(r.Body).Decode(&req); err != nil {

		http.Error(
			w,
			"datos inválidos",
			http.StatusBadRequest,
		)

		return
	}

	user, err :=
		h.UserRepo.GetByUsername(
			req.Username,
		)

	if err != nil {

		http.Error(
			w,
			"usuario o contraseña incorrectos",
			http.StatusUnauthorized,
		)

		return
	}

	if err :=
		bcrypt.CompareHashAndPassword(

			[]byte(
				user.PasswordHash,
			),

			[]byte(
				req.Password,
			),
		); err != nil {

		http.Error(
			w,
			"usuario o contraseña incorrectos",
			http.StatusUnauthorized,
		)

		return
	}

	token :=
		jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			jwt.MapClaims{

				"sub": user.ID,

				"username": user.Username,

				"exp": time.Now().
					Add(24 * time.Hour).
					Unix(),
			},
		)

	tokenStr, err :=
		token.SignedString(
			h.JWTKey,
		)

	if err != nil {

		http.Error(
			w,
			"error generando token",
			http.StatusInternalServerError,
		)

		return
	}

	http.SetCookie(
		w,
		&http.Cookie{

			Name: "token",

			Value: tokenStr,

			Path: "/",

			HttpOnly: true,

			SameSite: http.SameSiteLaxMode,

			// Cambiar a true
			// cuando la aplicación
			// use HTTPS.
			Secure: false,

			MaxAge: 60 * 60 * 24,
		},
	)

	w.WriteHeader(
		http.StatusOK,
	)
}

// LogoutHandler elimina la sesión
// del usuario borrando la cookie.
func (h AuthHandler) LogoutHandler(
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

			Secure: false,

			MaxAge: -1,
		},
	)

	if err :=
		json.NewEncoder(w).Encode(

			map[string]bool{

				"success": true,
			},
		); err != nil {

		http.Error(

			w,

			"error generando respuesta",

			http.StatusInternalServerError,
		)

	}
}
