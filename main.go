package main

import (
	"fmt"
	"net/http"
	"os"
	api "web-blog/handlers"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found")
	}
	if _, err := os.Stat("articles"); os.IsNotExist(err) {
		os.Mkdir("articles", 0755)
	}

}

func main() {
	http.HandleFunc("/", api.HomeHandler)
	http.HandleFunc("/articles/", api.ArticleHandler)
	http.HandleFunc("/search", api.SearchHandler)
	http.HandleFunc("/dashboard", api.DashboardArticleWithAuthI())
	http.HandleFunc("/new", api.CreateArticleWithAuthI())
	http.HandleFunc("/edit/", api.UpdateArticleWithAuthI())
	http.HandleFunc("/delete/", api.DeleteArticleWithAuthI())

	http.HandleFunc("/login", api.LoginHandler)
	http.HandleFunc("/logout", api.LogoutHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
