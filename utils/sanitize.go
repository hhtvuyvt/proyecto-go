package utils

import (
	"html"
	"strings"

	"github.com/hhtvuyvt/proyecto-go/models"
)

// SanitizeBook limpia y sanea los campos de texto de un libro.
func SanitizeBook(b *models.Book) {
	if b == nil {
		return
	}

	b.Title = html.EscapeString(strings.TrimSpace(b.Title))
	b.Author = html.EscapeString(strings.TrimSpace(b.Author))
	b.ISBN = html.EscapeString(strings.TrimSpace(b.ISBN))
	b.Image = strings.TrimSpace(b.Image)
}
