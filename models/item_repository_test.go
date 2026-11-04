package models

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite" // Necesario para que funcione la DB en memoria
)

// setupTestDB crea una base de datos en memoria y aplica solo el esquema
// necesario para probar ItemRepository, evitando ciclos de importación.
func setupTestDB(t *testing.T) *ItemRepository {
	// ":memory:" crea una base de datos temporal que se destruye al cerrar
	database, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Error abriendo DB de prueba: %v", err)
	}

	// Aseguramos que la DB se cierre al finalizar el test
	t.Cleanup(func() {
		_ = database.Close()
	})

	// Creamos manualmente la tabla 'items' para esta prueba específica
	schema := `
	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		category TEXT NOT NULL,
		code TEXT NOT NULL,
		image TEXT NOT NULL
	);
	`
	if _, err := database.Exec(schema); err != nil {
		t.Fatalf("Error creando esquema de prueba: %v", err)
	}

	return &ItemRepository{DB: database}
}

func TestItemRepository_CRUD(t *testing.T) {
	repo := setupTestDB(t)

	// Datos de prueba
	item := &Item{
		Name:     "Laptop Gamer",
		Category: "Electrónica",
		Code:     "SKU-12345",
		Image:    "laptop.jpg",
	}

	// ==========================================
	// 1. Probar Create
	// ==========================================
	err := repo.Create(item)
	if err != nil {
		t.Fatalf("Error creando item: %v", err)
	}
	if item.ID == 0 {
		t.Error("Se esperaba que el ID se generara automáticamente")
	}

	// ==========================================
	// 2. Probar GetByID
	// ==========================================
	fetchedItem, err := repo.GetByID(item.ID)
	if err != nil {
		t.Fatalf("Error obteniendo item por ID: %v", err)
	}
	if fetchedItem.Name != item.Name {
		t.Errorf("Esperaba nombre '%s', obtuvo '%s'", item.Name, fetchedItem.Name)
	}
	if fetchedItem.Code != item.Code {
		t.Errorf("Esperaba código '%s', obtuvo '%s'", item.Code, fetchedItem.Code)
	}

	// ==========================================
	// 3. Probar GetAll
	// ==========================================
	items, err := repo.GetAll()
	if err != nil {
		t.Fatalf("Error obteniendo todos los items: %v", err)
	}
	if len(items) != 1 {
		t.Errorf("Esperaba 1 item en la base de datos, obtuvo %d", len(items))
	}

	// ==========================================
	// 4. Probar Update
	// ==========================================
	item.Name = "Laptop Gamer Pro"
	item.Category = "Computación Avanzada"

	err = repo.Update(item)
	if err != nil {
		t.Fatalf("Error actualizando item: %v", err)
	}

	updatedItem, err := repo.GetByID(item.ID)
	if err != nil {
		t.Fatalf("Error obteniendo item actualizado: %v", err)
	}
	if updatedItem.Name != "Laptop Gamer Pro" {
		t.Errorf("El nombre no se actualizó correctamente. Esperaba 'Laptop Gamer Pro', obtuvo '%s'", updatedItem.Name)
	}

	// ==========================================
	// 5. Probar Delete
	// ==========================================
	err = repo.Delete(item.ID)
	if err != nil {
		t.Fatalf("Error eliminando item: %v", err)
	}

	// Verificar que ya no existe
	_, err = repo.GetByID(item.ID)
	if err == nil {
		t.Error("Se esperaba un error (sql.ErrNoRows) al obtener un item eliminado")
	}
	if err != sql.ErrNoRows {
		t.Errorf("Se esperaba sql.ErrNoRows, obtuvo: %v", err)
	}
}

// TestItemRepository_GetByID_NotFound prueba el comportamiento cuando el ID no existe
func TestItemRepository_GetByID_NotFound(t *testing.T) {
	repo := setupTestDB(t)

	_, err := repo.GetByID(999)
	if err == nil {
		t.Error("Se esperaba un error al buscar un ID inexistente")
	}
	if err != sql.ErrNoRows {
		t.Errorf("Se esperaba sql.ErrNoRows, obtuvo: %v", err)
	}
}
