package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// UploadResponse representa la ruta pública
// de la imagen subida.
type UploadResponse struct {
	Path string `json:"path"`
}

// UploadImage recibe una imagen,
// la guarda en uploads/
// y devuelve su ubicación.
func UploadImage(
	w http.ResponseWriter,
	r *http.Request,
) {

	if err := os.MkdirAll(
		"uploads",
		0755,
	); err != nil {

		http.Error(
			w,
			"error creando carpeta",
			http.StatusInternalServerError,
		)

		return
	}

	file, header, err :=
		r.FormFile("image")

	if err != nil {

		http.Error(
			w,
			"archivo inválido",
			http.StatusBadRequest,
		)

		return
	}

	defer func() {

		if err := file.Close(); err != nil {

			return

		}

	}()

	dstPath :=
		filepath.Join(
			"uploads",
			header.Filename,
		)

	dst, err :=
		os.Create(dstPath)

	if err != nil {

		http.Error(
			w,
			"error creando archivo",
			http.StatusInternalServerError,
		)

		return
	}

	_, copyErr :=
		io.Copy(
			dst,
			file,
		)

	closeErr :=
		dst.Close()

	if copyErr != nil {

		http.Error(
			w,
			"error guardando archivo",
			http.StatusInternalServerError,
		)

		return
	}

	if closeErr != nil {

		http.Error(
			w,
			"error cerrando archivo",
			http.StatusInternalServerError,
		)

		return
	}

	response :=
		UploadResponse{
			Path: "/uploads/" + header.Filename,
		}

	if err :=
		json.NewEncoder(w).Encode(response); err != nil {

		http.Error(
			w,
			"error respondiendo",
			http.StatusInternalServerError,
		)

		return
	}

}
