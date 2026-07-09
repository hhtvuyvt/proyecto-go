package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/hhtvuyvt/proyecto-go/models"
)

type fakeUserRepo struct{}

func (fakeUserRepo) GetByUsername(
	_ string,
) (models.User, error) {

	hash, _ :=
		bcrypt.GenerateFromPassword(

			[]byte("admin"),

			bcrypt.DefaultCost,
		)

	return models.User{

		ID: 1,

		Username: "admin",

		PasswordHash: string(hash),
	}, nil
}

func (fakeUserRepo) Create(
	_ *models.User,
) error {

	return nil

}

func TestLoginHandler(
	t *testing.T,
) {

	handler :=
		AuthHandler{

			UserRepo: fakeUserRepo{},

			JWTKey: []byte("secret"),
		}

	body, _ :=
		json.Marshal(

			LoginRequest{

				Username: "admin",

				Password: "admin",
			},
		)

	req :=
		httptest.NewRequest(

			http.MethodPost,

			"/api/login",

			bytes.NewReader(body),
		)

	rec :=
		httptest.NewRecorder()

	handler.LoginHandler(

		rec,

		req,
	)

	if rec.Code != http.StatusOK {

		t.Fatalf(

			"esperaba 200, obtuvo %d",

			rec.Code,
		)

	}

	cookies :=
		rec.Result().Cookies()

	if len(cookies) == 0 {

		t.Fatal(
			"no se creó ninguna cookie",
		)

	}

	if cookies[0].Name != "token" {

		t.Fatal(
			"la cookie no se llama token",
		)

	}

	if cookies[0].Value == "" {

		t.Fatal(
			"la cookie está vacía",
		)

	}

}

func TestLogoutHandler(
	t *testing.T,
) {

	handler := AuthHandler{}

	req :=
		httptest.NewRequest(
			http.MethodPost,
			"/api/logout",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	handler.LogoutHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {

		t.Fatalf(
			"esperaba 200, obtuvo %d",
			rec.Code,
		)

	}

	cookies :=
		rec.Result().Cookies()

	if len(cookies) == 0 {

		t.Fatal(
			"no se envió ninguna cookie",
		)

	}

	cookie :=
		cookies[0]

	if cookie.Name != "token" {

		t.Fatal(
			"la cookie debería llamarse token",
		)

	}

	if cookie.MaxAge != -1 {

		t.Fatal(
			"la cookie no fue eliminada",
		)

	}

}
