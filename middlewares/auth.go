package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	// ClaimsKey es la clave utilizada para almacenar los claims de JWT en el contexto.
	ClaimsKey contextKey = "claims"
)

// GetClaims extrae los claims del JWT almacenados en el contexto de la petición.
func GetClaims(r *http.Request) (jwt.MapClaims, bool) {
	claims, ok := r.Context().Value(ClaimsKey).(jwt.MapClaims)
	return claims, ok
}

// AuthMiddleware válida tokens JWT que viajan en cookies.
//
// La clave se recibe desde fuera para:
// - evitar dependencia de variables globales
// - facilitar tests
// - separar configuración de lógica
func AuthMiddleware(
	jwtKey []byte,
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			cookie, err :=
				r.Cookie(
					"token",
				)

			if err != nil {

				http.Error(
					w,
					"token requerido",
					http.StatusUnauthorized,
				)

				return
			}

			tokenStr :=
				cookie.Value

			token, err :=
				jwt.Parse(
					tokenStr,
					func(
						token *jwt.Token,
					) (interface{}, error) {

						// Evita aceptar algoritmos
						// diferentes al esperado.
						if _, ok :=
							token.Method.(*jwt.SigningMethodHMAC); !ok {

							return nil,
								fmt.Errorf(
									"algoritmo inválido",
								)
						}

						return jwtKey, nil
					},
				)

			if err != nil ||
				!token.Valid {

				http.Error(
					w,
					"token inválido",
					http.StatusUnauthorized,
				)

				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				ctx := context.WithValue(r.Context(), ClaimsKey, claims)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(
				w,
				r,
			)
		},
	)
}
