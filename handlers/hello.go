package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>¡Hola desde Go web!</h1><p>Ruta: %s</p>", r.URL.Path)
}
