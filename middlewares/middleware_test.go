package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggerMiddleware(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := LoggerMiddleware(nextHandler)

	req := httptest.NewRequest(http.MethodGet, "/test-logger", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba status 200, obtuvo %d", rec.Code)
	}
}

func TestRecoverMiddleware(t *testing.T) {
	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("pánico simulado para prueba")
	})

	handler := RecoverMiddleware(panicHandler)

	req := httptest.NewRequest(http.MethodGet, "/test-panic", nil)
	rec := httptest.NewRecorder()

	// El middleware debe interceptar el pánico internamente y responder con 500
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("esperaba status 500 tras el pánico, obtuvo %d", rec.Code)
	}
}
