package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestLoginHandler_GET(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	rr := httptest.NewRecorder()

	LoginHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "<form") {
		t.Fatal("expected login form in response body")
	}
}

func TestLoginHandler_POST_Success(t *testing.T) {
	os.Setenv("ADMIN_USERNAME", "admin")
	os.Setenv("ADMIN_PASSWORD", "123")

	form := "username=admin&password=123"
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	LoginHandler(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Fatalf("expected 303 redirect, got %d", rr.Code)
	}

	location := rr.Header().Get("Location")
	if location != "/dashboard" {
		t.Fatalf("expected redirect to /dashboard, got %s", location)
	}

	cookies := rr.Result().Cookies()
	if len(cookies) == 0 || cookies[0].Name != "auth_token" {
		t.Fatal("auth_token cookie not set")
	}
}

func TestLoginHandler_POST_Fail(t *testing.T) {
	os.Setenv("ADMIN_USERNAME", "admin")
	os.Setenv("ADMIN_PASSWORD", "123")

	form := "username=admin&password=wrongpass"
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	os.MkdirAll("templates", os.ModePerm)
	os.WriteFile("templates/login_error.html", []byte("LOGIN ERROR"), 0644)
	defer os.RemoveAll("templates")

	LoginHandler(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "LOGIN ERROR") {
		t.Fatal("expected login error template in response")
	}
}

func TestLogoutHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/logout", nil)
	rr := httptest.NewRecorder()

	LogoutHandler(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Fatalf("expected 303 redirect, got %d", rr.Code)
	}

	location := rr.Header().Get("Location")
	if location != "/login" {
		t.Fatalf("expected redirect to /login, got %s", location)
	}

	cookies := rr.Result().Cookies()
	if len(cookies) == 0 || cookies[0].Value != "" || cookies[0].MaxAge != -1 {
		t.Fatal("auth_token cookie not cleared")
	}
}
