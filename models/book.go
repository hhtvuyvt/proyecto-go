// Paquete models define las estructuras de datos de la aplicación.
package models

// Book representa un libro dentro del sistema.
type Book struct {
	ID     int64  `json:"id"`     // ID autoincremental
	Title  string `json:"title"`  // Título del libro
	Author string `json:"author"` // Autor del libro
	ISBN   string `json:"isbn"`   // Código ISBN
	Image  string `json:"image"`  // URL o nombre de la portada
}
