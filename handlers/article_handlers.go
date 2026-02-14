package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"web-blog/model"
)

func parseTemplates(templateName string) *template.Template {
	tmpl, err := template.ParseFiles("templates/catalog.html")
	if err != nil {
		panic(fmt.Sprintf("error parsing template %s, %v", templateName, err))
	}

	return tmpl
}

// CRUD
// Create
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var a model.Article
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, `{"error":"invalid JSON"}`, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	//Validation
	if a.Title == "" || len(a.Title) > 100 {
		http.Error(w, `{"error":"title required and  <100 chars"}`, http.StatusBadRequest)
		return
	}

	if a.Content == "" {
		http.Error(w, `{"error":"content required"}`, http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", a.Published); err != nil {
		http.Error(w, `{"error":"invalid date format YYYY-MM-DD"}`, http.StatusBadRequest)
		return
	}

	files, _ := os.ReadDir("articles")
	a.ID = len(files) + 1

	filePath := fmt.Sprintf("articles/article%d.json", a.ID)
	file, _ := os.Create(filePath)

	defer file.Close()
	json.NewEncoder(file).Encode(a)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
}

// Read all
func GetArticles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	files, _ := os.ReadDir("articles")
	var articles []model.Article

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".json" {
			continue
		}
		data, _ := os.ReadFile(filepath.Join("articles", f.Name()))
		var art model.Article
		json.Unmarshal(data, &art)
		articles = append(articles, art)
	}

	json.NewEncoder(w).Encode(articles)
}

// getArticle
// Update
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/articles"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}

	var a model.Article
	err = json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, `{"error":"invalid JSON"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	a.ID = id

	filePath := fmt.Sprintf("articles/article%d.json", a.ID)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		return
	}

	file, _ := os.Create(filePath)
	defer file.Close()
	json.NewEncoder(file).Encode(a)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)
}

// Delete
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/articles/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}

	filePath := fmt.Sprintf("articles/article%d.json", id)
	if err := os.Remove(filePath); err != nil {
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
