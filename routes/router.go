package routes

import (
	"net/http"

	"github.com/hhtvuyvt/proyecto-go/handlers"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HelloHandler)

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
