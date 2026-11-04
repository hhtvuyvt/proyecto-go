package utils

import (
	"html"
	"strings"

	"github.com/hhtvuyvt/proyecto-go/models"
)

// SanitizeItem limpia y sanea los campos de texto de un item (previene XSS).
func SanitizeItem(item *models.Item) {
	if item == nil {
		return
	}

	item.Name = html.EscapeString(strings.TrimSpace(item.Name))
	item.Category = html.EscapeString(strings.TrimSpace(item.Category))
	item.Code = html.EscapeString(strings.TrimSpace(item.Code))
	item.Image = strings.TrimSpace(item.Image)
}
