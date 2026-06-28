package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func createTestToken(t *testing.T) string {

	t.Helper()

	secret :=
		[]byte("test-secret")

	token :=
		jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			jwt.MapClaims{
				"user": "admin",
			},
		)

	signed, err :=
		token.SignedString(secret)

	if err != nil {
		t.Fatal(err)
	}

	return signed
}

func TestLoginHandler(t *testing.T) {

	t.Setenv(
		"JWT_SECRET",
		"test-secret",
	)

	req :=
		httptest.NewRequest(
			http.MethodPost,
			"/api/login",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	LoginHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {

		t.Fatalf(
			"esperado 200 recibido %d",
			rec.Code,
		)
	}

	var response map[string]string

	err :=
		json.NewDecoder(
			rec.Body,
		).Decode(&response)

	if err != nil {
		t.Fatal(err)
	}

	if response["token"] == "" {

		t.Fatal(
			"token vacío",
		)
	}

}

func TestAuthMiddlewareSinAuthorization(
	t *testing.T,
) {

	t.Setenv(
		"JWT_SECRET",
		"test-secret",
	)

	called :=
		false

	next :=
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {

				called = true
			},
		)

	handler :=
		AuthMiddleware(next)

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/books",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	handler.ServeHTTP(
		rec,
		req,
	)

	if rec.Code != http.StatusUnauthorized {

		t.Fatalf(
			"esperado 401 recibido %d",
			rec.Code,
		)
	}

	if called {
		t.Fatal(
			"pasó sin token",
		)
	}

}

func TestAuthMiddlewareTokenInvalido(
	t *testing.T,
) {

	t.Setenv(
		"JWT_SECRET",
		"test-secret",
	)

	handler :=
		AuthMiddleware(
			http.HandlerFunc(
				func(
					w http.ResponseWriter,
					r *http.Request,
				) {

					t.Fatal(
						"no debería ejecutarse",
					)
				},
			),
		)

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/books",
			nil,
		)

	req.Header.Set(
		"Authorization",
		"Bearer falso",
	)

	rec :=
		httptest.NewRecorder()

	handler.ServeHTTP(
		rec,
		req,
	)

	if rec.Code != http.StatusUnauthorized {

		t.Fatalf(
			"esperado 401 recibido %d",
			rec.Code,
		)
	}

}

func TestAuthMiddlewareTokenValido(
	t *testing.T,
) {

	t.Setenv(
		"JWT_SECRET",
		"test-secret",
	)

	passed :=
		false

	handler :=
		AuthMiddleware(
			http.HandlerFunc(
				func(
					w http.ResponseWriter,
					r *http.Request,
				) {

					passed = true

					w.WriteHeader(
						http.StatusOK,
					)

				},
			),
		)

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/books/1",
			nil,
		)

	req.Header.Set(
		"Authorization",
		"Bearer "+createTestToken(t),
	)

	rec :=
		httptest.NewRecorder()

	handler.ServeHTTP(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {

		t.Fatalf(
			"esperado 200 recibido %d",
			rec.Code,
		)
	}

	if !passed {

		t.Fatal(
			"middleware bloqueó token válido",
		)
	}

}

func TestAuthMiddlewareFormatoIncorrecto(
	t *testing.T,
) {

	t.Setenv(
		"JWT_SECRET",
		"test-secret",
	)

	handler :=
		AuthMiddleware(
			http.HandlerFunc(
				func(
					w http.ResponseWriter,
					r *http.Request,
				) {

					t.Fatal(
						"no debería pasar",
					)
				},
			),
		)

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/",
			nil,
		)

	req.Header.Set(
		"Authorization",
		"Basic token",
	)

	rec :=
		httptest.NewRecorder()

	handler.ServeHTTP(
		rec,
		req,
	)

	if rec.Code != http.StatusUnauthorized {

		t.Fatalf(
			"esperado 401 recibido %d",
			rec.Code,
		)
	}

	_ = strings.Builder{}
	_ = os.Getenv
}
