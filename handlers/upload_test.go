package handlers

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func createMultipartImage(
	t *testing.T,
) (*http.Request, string) {

	t.Helper()

	body :=
		&bytes.Buffer{}

	writer :=
		multipart.NewWriter(body)

	part, err :=
		writer.CreateFormFile(
			"image",
			"test.jpg",
		)

	if err != nil {
		t.Fatal(err)
	}

	_, err =
		part.Write(
			[]byte(
				"fake image content",
			),
		)

	if err != nil {
		t.Fatal(err)
	}

	if err :=
		writer.Close(); err != nil {

		t.Fatal(err)
	}

	req :=
		httptest.NewRequest(
			http.MethodPost,
			"/api/upload",
			body,
		)

	req.Header.Set(
		"Content-Type",
		writer.FormDataContentType(),
	)

	return req, "test.jpg"
}

func TestUploadImageSuccess(
	t *testing.T,
) {

	_ = os.RemoveAll("uploads")

	defer func() {
		_ = os.RemoveAll("uploads")
	}()

	req, filename :=
		createMultipartImage(t)

	rec :=
		httptest.NewRecorder()

	UploadImage(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {

		t.Fatalf(
			"esperado 200 recibido %d",
			rec.Code,
		)
	}

	var response UploadResponse

	err :=
		json.NewDecoder(
			rec.Body,
		).Decode(
			&response,
		)

	if err != nil {
		t.Fatal(err)
	}

	expected :=
		"/uploads/" + filename

	if response.Path != expected {

		t.Fatalf(
			"ruta incorrecta esperada %s recibida %s",
			expected,
			response.Path,
		)
	}

	filePath :=
		filepath.Join(
			"uploads",
			filename,
		)

	if _, err := os.Stat(filePath); err != nil {

		t.Fatalf(
			"archivo no creado: %v",
			err,
		)
	}

}

func TestUploadImageSinArchivo(
	t *testing.T,
) {

	_ = os.RemoveAll("uploads")

	defer func() {
		_ = os.RemoveAll("uploads")
	}()

	req :=
		httptest.NewRequest(
			http.MethodPost,
			"/api/upload",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	UploadImage(
		rec,
		req,
	)

	if rec.Code != http.StatusBadRequest {

		t.Fatalf(
			"esperado 400 recibido %d",
			rec.Code,
		)
	}

}

func TestUploadImageMultipartVacio(
	t *testing.T,
) {

	body :=
		&bytes.Buffer{}

	writer :=
		multipart.NewWriter(body)

	if err :=
		writer.Close(); err != nil {

		t.Fatal(err)
	}

	req :=
		httptest.NewRequest(
			http.MethodPost,
			"/api/upload",
			body,
		)

	req.Header.Set(
		"Content-Type",
		writer.FormDataContentType(),
	)

	rec :=
		httptest.NewRecorder()

	UploadImage(
		rec,
		req,
	)

	if rec.Code != http.StatusBadRequest {

		t.Fatalf(
			"esperado 400 recibido %d",
			rec.Code,
		)
	}

}
