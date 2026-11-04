package utils

import (
	"testing"

	"github.com/hhtvuyvt/proyecto-go/models"
)

func TestSanitizeItem(t *testing.T) {
	item := &models.Item{
		Name:     "  <script>alert('xss')</script> Laptop  ",
		Category: "  Electrónica  ",
		Code:     "  SKU-123  ",
		Image:    "  image.jpg  ",
	}

	SanitizeItem(item)

	// Verificamos que el HTML fue escapado y los espacios eliminados
	if item.Name == "<script>alert('xss')</script> Laptop" {
		t.Errorf("El nombre no fue saneado (XSS no prevenido): %s", item.Name)
	}

	if item.Category != "Electrónica" {
		t.Errorf("La categoría no fue saneada correctamente: %s", item.Category)
	}

	if item.Code != "SKU-123" {
		t.Errorf("El código no fue saneado correctamente: %s", item.Code)
	}

	if item.Image != "image.jpg" {
		t.Errorf("La imagen no fue saneada correctamente: %s", item.Image)
	}
}

func TestSanitizeItem_Nil(t *testing.T) {
	// No debería entrar en pánico si se pasa nil
	SanitizeItem(nil)
}
