// Paquete utils contiene utilidades comunes usadas en múltiples partes del proyecto.
package utils

import (
	"html"
	"strings"

	"github.com/hhtvuyvt/proyecto-go/models"
)

// SanitizeBook limpia campos de texto en un libro para evitar inyecciones HTML o entradas maliciosas.
// Reemplaza caracteres especiales por sus entidades HTML y recorta espacios en blanco.
func SanitizeBook(b *models.Book) {
	b.Title = html.EscapeString(strings.TrimSpace(b.Title))
	b.Author = html.EscapeString(strings.TrimSpace(b.Author))
	b.ISBN = html.EscapeString(strings.TrimSpace(b.ISBN))
	b.Image = html.EscapeString(strings.TrimSpace(b.Image))
}
