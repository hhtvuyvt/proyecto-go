package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hhtvuyvt/proyecto-go/internal/db"
	"github.com/hhtvuyvt/proyecto-go/models"
)

func createTestRouter(t *testing.T) http.Handler {
	t.Helper()

	t.Setenv("ADMIN_USERNAME", "admin")
	t.Setenv("ADMIN_PASSWORD", "admin123")

	testDB, err := db.Open(":memory:")
	if err != nil {
		t.Fatalf("no se pudo abrir la base de datos de prueba: %v", err)
	}

	// 1. AÑADIDO: ItemRepo es obligatorio en el nuevo RouterConfig
	itemRepo := &models.ItemRepository{DB: testDB}
	detailRepo := &models.DetailRepository{DB: testDB}
	userRepo := &models.UserRepository{DB: testDB}

	return Router(
		RouterConfig{
			ItemRepo:   itemRepo,   // <-- AÑADIDO
			DetailRepo: detailRepo, // <-- CORREGIDO (antes ChapterRepo)
			UserRepo:   userRepo,
			JWTKey:     []byte("test-secret"),
		},
	)
}

func TestRouterRedirect(t *testing.T) {
	router := createTestRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusTemporaryRedirect {
		t.Fatalf("esperaba redirect %d, obtuvo %d", http.StatusTemporaryRedirect, rec.Code)
	}
}

func TestRouterStatic(t *testing.T) {
	router := createTestRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/static/index.html", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code == http.StatusNotFound {
		t.Fatal("ruta static no configurada")
	}
}

func TestRouterLoginRoute(t *testing.T) {
	router := createTestRouter(t)

	req := httptest.NewRequest(http.MethodPost, "/api/login", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code == http.StatusNotFound {
		t.Fatal("ruta /api/login no existe")
	}
}

// 2. ACTUALIZADO: De books a items
func TestRouterItemsPublicRoute(t *testing.T) {
	router := createTestRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/api/items", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code == http.StatusNotFound {
		t.Fatal("ruta /api/items no existe")
	}
}

// 3. ACTUALIZADO: URL a /api/items/1
func TestRouterProtectedItems(t *testing.T) {
	router := createTestRouter(t)

	req := httptest.NewRequest(http.MethodDelete, "/api/items/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	// Sin JWT debe bloquear con 401
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401 sin token, obtuvo %d", rec.Code)
	}
}

// 4. ACTUALIZADO: URL a /api/items
func TestRouterItemsPostRequiresAuth(t *testing.T) {
	router := createTestRouter(t)

	req := httptest.NewRequest(http.MethodPost, "/api/items", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401 sin token, obtuvo %d", rec.Code)
	}
}

// 5. ACTUALIZADO: URL a /api/items/1/details
func TestRouterDetailsPostRequiresAuth(t *testing.T) {
	router := createTestRouter(t)

	req := httptest.NewRequest(http.MethodPost, "/api/items/1/details", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401 sin token, obtuvo %d", rec.Code)
	}
}

func TestRouterMeRequiresAuth(t *testing.T) {
	router := createTestRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/api/me", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401 sin token, obtuvo %d", rec.Code)
	}
}
