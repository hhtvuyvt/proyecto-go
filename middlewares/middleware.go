package middlewares

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// sanitizeLog sanitizes user input to prevent log injection (CWE-117).
func sanitizeLog(str string) string {
	replacer := strings.NewReplacer("\n", "", "\r", "")
	return replacer.Replace(str)
}

func LoggerMiddleware(
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			start := time.Now()
			safePath := sanitizeLog(r.URL.Path)
			safeMethod := sanitizeLog(r.Method)

			log.Printf(
				"⏳ %s %s",
				safeMethod,
				safePath,
			)

			next.ServeHTTP(
				w,
				r,
			)

			log.Printf(
				"✅ %s %s (%s)",
				safeMethod,
				safePath,
				time.Since(start),
			)
		},
	)
}

func RecoverMiddleware(
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			defer func() {

				if err := recover(); err != nil {

					log.Printf(
						"💥 Panic: %v",
						err,
					)

					http.Error(
						w,
						"Error interno del servidor",
						http.StatusInternalServerError,
					)
				}

			}()

			next.ServeHTTP(
				w,
				r,
			)
		},
	)
}
