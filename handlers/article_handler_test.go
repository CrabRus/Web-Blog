package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"web-blog/handlers/middleware"
	"web-blog/model"
)

// ---------- HELPERS ----------

func setupTestEnv(t *testing.T) string {
	dir := t.TempDir()

	articlesDir := filepath.Join(dir, "articles")
	os.MkdirAll(articlesDir, os.ModePerm)

	templatesDir := filepath.Join(dir, "templates")
	os.MkdirAll(templatesDir, os.ModePerm)

	os.WriteFile(filepath.Join(templatesDir, "home.html"),
		[]byte("HOME"), 0644)

	os.WriteFile(filepath.Join(templatesDir, "dashboard.html"),
		[]byte("DASHBOARD"), 0644)

	os.WriteFile(filepath.Join(templatesDir, "newArticle.html"),
		[]byte("NEW"), 0644)

	os.WriteFile(filepath.Join(templatesDir, "updateArticle.html"),
		[]byte("UPDATE"), 0644)

	oldWd, _ := os.Getwd()
	os.Chdir(dir)

	t.Cleanup(func() {
		os.Chdir(oldWd)
	})

	return articlesDir
}

func addAdmin(req *http.Request) *http.Request {
	claims := &middleware.Claims{
		Username: "admin",
		Role:     "admin",
	}
	ctx := context.WithValue(req.Context(), "user", claims)
	return req.WithContext(ctx)
}

// ---------- CREATE ----------

func TestCreateArticle_Success(t *testing.T) {

	articlesDir := setupTestEnv(t)

	form := "title=Hello&content=World&published=2025-02-17"
	req := httptest.NewRequest(http.MethodPost, "/new", bytes.NewBufferString(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req = addAdmin(req)

	rr := httptest.NewRecorder()

	createArticle(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Fatalf("expected 303, got %d", rr.Code)
	}

	files, _ := os.ReadDir(articlesDir)
	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}
}

// ---------- CREATE UNAUTHORIZED ----------

func TestCreateArticle_Unauthorized(t *testing.T) {

	setupTestEnv(t)

	form := "title=Hello&content=World&published=2025-02-17"
	req := httptest.NewRequest(http.MethodPost, "/new", bytes.NewBufferString(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	createArticle(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}
}

// ---------- GET ARTICLE ----------

func TestGetArticleByID(t *testing.T) {

	articlesDir := setupTestEnv(t)

	article := model.Article{
		ID:    1,
		Title: "Test",
	}

	data, _ := json.Marshal(article)
	os.WriteFile(filepath.Join(articlesDir, "article1.json"), data, 0644)

	a := GetArticleByID(1)

	if a == nil {
		t.Fatal("expected article, got nil")
	}

	if a.Title != "Test" {
		t.Fatalf("expected Test, got %s", a.Title)
	}
}

// ---------- UPDATE ----------

func TestUpdateArticle(t *testing.T) {

	articlesDir := setupTestEnv(t)

	os.WriteFile(filepath.Join(articlesDir, "article1.json"),
		[]byte(`{"id":1}`), 0644)

	form := "title=Updated&content=New&published=2025-02-18"
	req := httptest.NewRequest(http.MethodPut, "/edit/1", bytes.NewBufferString(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	updateArticle(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}

// ---------- UPDATE INVALID ID ----------

func TestUpdateArticle_InvalidID(t *testing.T) {

	setupTestEnv(t)

	req := httptest.NewRequest(http.MethodPut, "/edit/abc", nil)
	rr := httptest.NewRecorder()

	updateArticle(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

// ---------- DELETE SUCCESS ----------

func TestDeleteArticle_Success(t *testing.T) {

	articlesDir := setupTestEnv(t)

	filePath := filepath.Join(articlesDir, "article1.json")
	os.WriteFile(filePath, []byte("test"), 0644)

	req := httptest.NewRequest(http.MethodDelete, "/delete/1", nil)
	rr := httptest.NewRecorder()

	deleteArticle(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		t.Fatal("file was not deleted")
	}
}

// ---------- DELETE NOT FOUND ----------

func TestDeleteArticle_NotFound(t *testing.T) {

	setupTestEnv(t)

	req := httptest.NewRequest(http.MethodDelete, "/delete/99", nil)
	rr := httptest.NewRecorder()

	deleteArticle(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", rr.Code)
	}
}

func TestCreateArticle_Get(t *testing.T) {
	setupTestEnv(t)

	req := httptest.NewRequest(http.MethodGet, "/new", nil)
	rr := httptest.NewRecorder()

	createArticle(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}

func TestCreateArticle_InvalidTitle(t *testing.T) {
	setupTestEnv(t)

	form := "title=&content=World&published=2025-02-17"
	req := httptest.NewRequest(http.MethodPost, "/new", bytes.NewBufferString(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req = addAdmin(req)

	rr := httptest.NewRecorder()

	createArticle(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestUpdateArticle_MethodNotAllowed(t *testing.T) {
	setupTestEnv(t)

	req := httptest.NewRequest(http.MethodPost, "/edit/1", nil)
	rr := httptest.NewRecorder()

	updateArticle(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", rr.Code)
	}
}

func TestDeleteArticle_MethodNotAllowed(t *testing.T) {
	setupTestEnv(t)

	req := httptest.NewRequest(http.MethodPost, "/delete/1", nil)
	rr := httptest.NewRecorder()

	deleteArticle(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", rr.Code)
	}
}

func TestHomeHandler(t *testing.T) {
	setupTestEnv(t)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	HomeHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
}
