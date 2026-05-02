package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// UploadImage permite subir imágenes al servidor.
func UploadImage(w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "archivo inválido", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dstPath := filepath.Join("uploads", header.Filename)

	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "error creando archivo", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "error guardando archivo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}