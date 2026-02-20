package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"web-blog/model"
)

func TestGetMaxArticleID(t *testing.T) {
	dir := t.TempDir()

	os.WriteFile(filepath.Join(dir, "article1.json"), []byte("{}"), 0644)
	os.WriteFile(filepath.Join(dir, "article5.json"), []byte("{}"), 0644)
	os.WriteFile(filepath.Join(dir, "article3.json"), []byte("{}"), 0644)

	maxID := GetMaxArticleID(dir)

	if maxID != 5 {
		t.Fatalf("expected 5, got %d", maxID)
	}
}

func TestCreateArticleByFilePath(t *testing.T) {

	dir := t.TempDir()

	article := model.Article{
		ID:        1,
		Title:     "Test",
		Content:   "Content",
		Published: "true",
		Author:    "admin",
	}

	err := CreateArticleByFilePath(dir, article)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	filePath := filepath.Join(dir, "article1.json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("file was not created")
	}
}

func TestCreateArticleContent(t *testing.T) {

	dir := t.TempDir()

	article := model.Article{
		ID:        2,
		Title:     "Hello",
		Content:   "World",
		Published: "true",
		Author:    "admin",
	}

	err := CreateArticleByFilePath(dir, article)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "article2.json"))
	if err != nil {
		t.Fatal(err)
	}

	var saved model.Article
	if err := json.Unmarshal(data, &saved); err != nil {
		t.Fatal(err)
	}

	if saved.Title != "Hello" {
		t.Fatalf("expected Hello, got %s", saved.Title)
	}
}

func TestParseTemplates_Success(t *testing.T) {
	dir := t.TempDir()

	templatesDir := filepath.Join(dir, "templates")
	os.MkdirAll(templatesDir, os.ModePerm)

	templateContent := `{{define "test"}}Hello{{end}}`
	templatePath := filepath.Join(templatesDir, "test.html")
	os.WriteFile(templatePath, []byte(templateContent), 0644)

	oldWd, _ := os.Getwd()
	defer os.Chdir(oldWd)

	os.Chdir(dir)

	tmpl := ParseTemplates("test.html")

	if tmpl == nil {
		t.Fatal("expected template, got nil")
	}
}

func TestParseTemplates_Panic(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic, but did not panic")
		}
	}()

	ParseTemplates("not_exists.html")
}
