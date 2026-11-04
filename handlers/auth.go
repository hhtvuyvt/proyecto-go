package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/hhtvuyvt/proyecto-go/middlewares"
	"github.com/hhtvuyvt/proyecto-go/models"
)

// AuthHandler maneja la autenticación de usuarios.
type AuthHandler struct {
	UserRepo models.UserRepositoryInterface
	JWTKey   []byte
}

// LoginRequest representa el cuerpo JSON recibido en el formulario de login.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler autentica al usuario y configura una cookie HttpOnly con el JWT.
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "datos inválidos", http.StatusBadRequest)
		return
	}

	user, err := h.UserRepo.GetByUsername(req.Username)
	if err != nil {
		http.Error(w, "usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, "usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString(h.JWTKey)
	if err != nil {
		http.Error(w, "error generando token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenStr,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		MaxAge:   60 * 60 * 24,
	})

	w.WriteHeader(http.StatusOK)
}

// LogoutHandler elimina la sesión del usuario borrando la cookie.
func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		MaxAge:   -1,
	})

	if err := json.NewEncoder(w).Encode(map[string]bool{"success": true}); err != nil {
		http.Error(w, "error generando respuesta", http.StatusInternalServerError)
	}
}

// MeHandler devuelve la información del usuario autenticado actual a partir del contexto.
func (h *AuthHandler) MeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	claims, ok := middlewares.GetClaims(r)
	if !ok {
		http.Error(w, "no autorizado", http.StatusUnauthorized)
		return
	}

	username, ok := claims["username"].(string)
	if !ok || username == "" {
		http.Error(w, "no autorizado", http.StatusUnauthorized)
		return
	}

	user, err := h.UserRepo.GetByUsername(username)
	if err != nil {
		http.Error(w, "usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "error generando respuesta", http.StatusInternalServerError)
		return
	}
}

// RegisterRequest representa el cuerpo JSON para crear un nuevo usuario.
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// Puedes agregar más campos aquí para SaaS, ej: Email, CompanyName, etc.
}

// RegisterHandler crea un nuevo usuario en la base de datos.
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "datos inválidos", http.StatusBadRequest)
		return
	}

	// Validación básica
	if req.Username == "" || req.Password == "" {
		http.Error(w, "usuario y contraseña son obligatorios", http.StatusBadRequest)
		return
	}

	// 1. Verificar si el usuario ya existe
	_, err := h.UserRepo.GetByUsername(req.Username)
	if err == nil {
		// Si no devuelve error, es que el usuario YA existe
		http.Error(w, "el usuario ya existe", http.StatusConflict)
		return
	}

	// 2. Hashear la contraseña (Costo 10 es el estándar recomendado)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "error interno del servidor", http.StatusInternalServerError)
		return
	}

	// 3. Crear el objeto usuario
	newUser := models.User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		// Role: "tenant", // Descomenta si tienes un campo de rol en tu modelo
	}

	// 4. Guardar en la base de datos
	// NOTA: Necesitarás agregar el método Create a tu UserRepositoryInterface
	err = h.UserRepo.Create(&newUser)
	if err != nil {
		http.Error(w, "error creando el usuario", http.StatusInternalServerError)
		return
	}

	// 5. Responder con éxito (Sin devolver la contraseña, por supuesto)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "usuario creado exitosamente",
		"username": newUser.Username,
	})
}
