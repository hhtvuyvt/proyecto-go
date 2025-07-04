package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "archivo demasiado grande", http.StatusBadRequest)
		return
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "no se recibió el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	contentType := header.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		http.Error(w, "solo se permiten imágenes", http.StatusBadRequest)
		return
	}

	filename := strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(header.Filename)
	path := filepath.Join("uploads", filename)

	dst, err := os.Create(path)
	if err != nil {
		http.Error(w, "no se pudo guardar el archivo", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	io.Copy(dst, file)

	url := "/uploads/" + filename
	json.NewEncoder(w).Encode(map[string]string{"url": url})
}
