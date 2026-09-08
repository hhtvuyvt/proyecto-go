package utils

import (
	"testing"

	"github.com/hhtvuyvt/proyecto-go/models"
)

func TestSanitizeBook(t *testing.T) {
	book := &models.Book{
		Title:  "   Golang Guide <script>alert('xss')</script>   ",
		Author: "   John 'The Creator' Doe   ",
		ISBN:   "   978-3-16-148410-0   ",
		Image:  "   /uploads/books/golang.png   ",
	}

	SanitizeBook(book)

	expectedTitle := "Golang Guide &lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;"
	if book.Title != expectedTitle {
		t.Errorf("título incorrecto.\nEsperado: %s\nObtenido: %s", expectedTitle, book.Title)
	}

	expectedAuthor := "John &#39;The Creator&#39; Doe"
	if book.Author != expectedAuthor {
		t.Errorf("autor incorrecto.\nEsperado: %s\nObtenido: %s", expectedAuthor, book.Author)
	}

	expectedISBN := "978-3-16-148410-0"
	if book.ISBN != expectedISBN {
		t.Errorf("ISBN incorrecto.\nEsperado: %s\nObtenido: %s", expectedISBN, book.ISBN)
	}

	expectedImage := "/uploads/books/golang.png"
	if book.Image != expectedImage {
		t.Errorf("imagen incorrecta.\nEsperado: %s\nObtenido: %s", expectedImage, book.Image)
	}
}

func TestSanitizeBookNil(t *testing.T) {
	// Verifica que no genere panic si se pasa un puntero nil
	SanitizeBook(nil)
}
