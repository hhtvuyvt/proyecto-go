package middlewares

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

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

			next.ServeHTTP(
				w,
				r,
			)
		},
	)
}
