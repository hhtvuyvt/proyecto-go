package utils

import (
	"regexp"
	"strings"

	"github.com/hhtvuyvt/proyecto-go/models"
)

func SanitizeBook(b *models.Book) {
	b.Title = normalize(strings.TrimSpace(b.Title))
	b.Author = normalize(strings.TrimSpace(b.Author))
	b.ISBN = strings.ReplaceAll(strings.TrimSpace(b.ISBN), "-", "")
	b.Image = strings.TrimSpace(b.Image)
}

func normalize(s string) string {
	return regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
}
