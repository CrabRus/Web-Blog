package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"html/template"
	"net/http"
	"os"
	"web-blog/handlers/middleware"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
	Success string `json:"success"`
}

func SetAuthCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		MaxAge:   24 * 60 * 60,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
}

func hashPassword(pass string) string {
	hash := sha256.Sum256([]byte(pass))
	return hex.EncodeToString(hash[:])
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
    <!DOCTYPE html>
    <html>
    <head>
       <title>Login</title>
    </head>
    <body>
       <h2>Admin Login</h2>
       <form method="POST" action = "/login">
          <input type="text" name = "username" placeholder= "Username" required>
          <input type="password" name = "password" placeholder= "Password" required>
          <button type = "submit">Login</button>
       </form>
    </body>
    </html>
    `))
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminUsername == "" {
		adminUsername = "admin"
	}

	if adminPassword == "" {
		adminPassword = "123"
	}

	if username != adminUsername || hashPassword(password) != hashPassword(adminPassword) {
		w.WriteHeader(http.StatusUnauthorized)
		tmpl := template.Must(template.ParseFiles("templates/login_error.html"))
		tmpl.Execute(w, nil)
		return
	}

	token, err := middleware.GenerateJWT(username)
	if err != nil {
		http.Error(w, `{"error":"failed generate token"}`, http.StatusInternalServerError)
		return
	}

	SetAuthCookie(w, token)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	middleware.ClearAuthCookie(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
