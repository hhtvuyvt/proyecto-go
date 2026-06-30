package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestLoginHandlerGeneraToken(
	t *testing.T,
) {

	err :=
		os.Setenv(
			"JWT_SECRET",
			"secret-test",
		)

	if err != nil {

		t.Fatal(
			"no se pudo configurar JWT_SECRET:",
			err,
		)

	}

	req :=
		httptest.NewRequest(
			http.MethodGet,
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
			"esperaba 200, recibido %d",
			rec.Code,
		)

	}

	if !strings.Contains(
		rec.Body.String(),
		"token",
	) {

		t.Fatal(
			"la respuesta no contiene token",
		)

	}
}
