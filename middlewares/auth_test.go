package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// genera un token JWT válido para tests
func generateTestToken(secret []byte) string {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
		"sub": "test-user",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString(secret)
	return tokenStr
}

func TestAuthMiddlewareWithoutToken(t *testing.T) {
	secret := []byte("test-secret")

	handler := AuthMiddleware(secret, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401 sin token, obtuve %d", rec.Code)
	}
}

func TestAuthMiddlewareWithInvalidToken(t *testing.T) {
	secret := []byte("test-secret")

	handler := AuthMiddleware(secret, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer token-invalido")

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401 con token inválido, obtuve %d", rec.Code)
	}
}

func TestAuthMiddlewareWithValidToken(t *testing.T) {
	secret := []byte("test-secret")
	token := generateTestToken(secret)

	handler := AuthMiddleware(secret, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200 con token válido, obtuve %d", rec.Code)
	}
}
