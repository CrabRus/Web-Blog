package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"web-blog/handlers/middleware"
	"web-blog/model"
	"web-blog/utils"
)

func parseTemplates(templateName string) *template.Template {
	tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s", templateName))
	if err != nil {
		panic(fmt.Sprintf("error parsing template %s, %v", templateName, err))
	}
	return tmpl
}

// CRUD
// Create
func createArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := parseTemplates("newArticle.html")
		tmpl.Execute(w, nil)
		return
	}

	user := middleware.GetUserFromContext(r)
	if user == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	published := r.FormValue("published")

	// Validation
	if err := utils.ValidateTitle(title); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	if err := utils.ValidateContent(content); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	if err := utils.ValidatePublished(published); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	files, _ := os.ReadDir("articles")
	maxID := 0

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			var id int
			fmt.Scanf(file.Name(), "post%d.json", &id)
			if id > maxID {
				maxID = id
			}
		}
	}

	var a model.Article
	a.Title = title
	a.Content = content
	a.Published = published
	a.Author = user.Username
	a.ID = maxID + 1

	filePath := fmt.Sprintf("articles/article%d.json", a.ID)
	file, _ := os.Create(filePath)

	defer file.Close()
	json.NewEncoder(file).Encode(a)

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func getArticles() []model.Article {
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
	return articles
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	articles := getArticles()
	tmpl := parseTemplates("home.html")
	tmpl.Execute(w, articles)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	articles := getArticles()
	tmpl := parseTemplates("dashboard.html")
	tmpl.Execute(w, articles)
}

// getArticle
func getArticleByID(id int) *model.Article {
	filePath := fmt.Sprintf("articles/article%d.json", id)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}
	var a model.Article
	json.Unmarshal(data, &a)
	return &a
}

// Update
func updateArticle(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/edit/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodGet {

		article := getArticleByID(id)
		if article == nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		tmpl := parseTemplates("updateArticle.html")
		tmpl.Execute(w, article)
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	published := r.FormValue("published")

	article := model.Article{
		ID:        id,
		Title:     title,
		Content:   content,
		Published: published,
	}

	filePath := fmt.Sprintf("articles/article%d.json", id)

	file, _ := os.Create(filePath)
	defer file.Close()

	json.NewEncoder(file).Encode(article)

	w.WriteHeader(http.StatusOK)
}

// Delete
func deleteArticle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/delete/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	filePath := fmt.Sprintf("articles/article%d.json", id)

	if err := os.Remove(filePath); err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CreateArticleWithAuthI() http.HandlerFunc {
	return middleware.CookieAuthMiddleware(createArticle)
}

func DashboardArticleWithAuthI() http.HandlerFunc {
	return middleware.CookieAuthMiddleware(dashboardHandler)
}
func UpdateArticleWithAuthI() http.HandlerFunc {
	return middleware.CookieAuthMiddleware(updateArticle)
}
func DeleteArticleWithAuthI() http.HandlerFunc {
	return middleware.CookieAuthMiddleware(deleteArticle)
}
