package models

// Item representa un recurso genérico en el sistema SaaS
// (puede ser un Producto en una tienda o una Habitación en un hotel).
type Item struct {
	ID       int64  `json:"id"`       // ID autoincremental
	Name     string `json:"name"`     // Nombre del producto / habitación
	Category string `json:"category"` // Categoría / Proveedor / Tipo
	Code     string `json:"code"`     // SKU / Código de barras / Número de habitación
	Image    string `json:"image"`    // URL o nombre de la imagen
}
