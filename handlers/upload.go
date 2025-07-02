package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "imagen no válida", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		http.Error(w, "formato de imagen no permitido", http.StatusBadRequest)
		return
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	path := filepath.Join("uploads", filename)

	out, err := os.Create(path)
	if err != nil {
		http.Error(w, "error guardando imagen", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		http.Error(w, "error escribiendo imagen", http.StatusInternalServerError)
		return
	}

	url := "/uploads/" + filename
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"url":"` + url + `"}`))
}
