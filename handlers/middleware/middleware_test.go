package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// ---------- HELPER ----------

func setupJWT() {
	os.Setenv("JWT_SECRET", "testsecret")
}

// ---------- JWTMiddleware ----------

func TestJWTMiddleware_MissingHeader(t *testing.T) {
	setupJWT()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handler := JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {})
	handler(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}
}

func TestJWTMiddleware_InvalidToken(t *testing.T) {
	setupJWT()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")

	rr := httptest.NewRecorder()

	handler := JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {})
	handler(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}
}

func TestJWTMiddleware_ValidToken(t *testing.T) {
	setupJWT()

	token, _ := GenerateJWT("admin")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()

	called := false
	handler := JWTMiddleware(func(w http.ResponseWriter, r *http.Request) {
		called = true

		user := GetUserFromContext(r)
		if user == nil || user.Username != "admin" {
			t.Fatal("user not found in context")
		}
	})

	handler(rr, req)

	if !called {
		t.Fatal("next handler not called")
	}
}

// ---------- CookieAuthMiddleware ----------

func TestCookieAuthMiddleware_NoCookie(t *testing.T) {
	setupJWT()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handler := CookieAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {})
	handler(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Fatalf("expected 303 redirect, got %d", rr.Code)
	}
}

func TestCookieAuthMiddleware_InvalidToken(t *testing.T) {
	setupJWT()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "auth_token",
		Value: "invalid",
	})

	rr := httptest.NewRecorder()

	handler := CookieAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {})
	handler(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Fatalf("expected 303 redirect, got %d", rr.Code)
	}
}

func TestCookieAuthMiddleware_ValidToken(t *testing.T) {
	setupJWT()

	token, _ := GenerateJWT("admin")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "auth_token",
		Value: token,
	})

	rr := httptest.NewRecorder()

	called := false

	handler := CookieAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		called = true
		user := GetUserFromContext(r)
		if user == nil {
			t.Fatal("user missing in context")
		}
	})

	handler(rr, req)

	if !called {
		t.Fatal("next not called")
	}
}

// ---------- AdminOnly ----------

func TestAdminOnly_Unauthorized(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handler := AdminOnly(func(w http.ResponseWriter, r *http.Request) {})
	handler(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rr.Code)
	}
}

func TestAdminOnly_Forbidden(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	ctx := context.WithValue(req.Context(), "user", &Claims{
		Username: "user",
		Role:     "user",
	})
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	handler := AdminOnly(func(w http.ResponseWriter, r *http.Request) {})
	handler(rr, req)

	if rr.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", rr.Code)
	}
}

func TestAdminOnly_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	ctx := context.WithValue(req.Context(), "user", &Claims{
		Username: "admin",
		Role:     "admin",
	})
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	called := false
	handler := AdminOnly(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	handler(rr, req)

	if !called {
		t.Fatal("next not called")
	}
}

// ---------- ClearAuthCookie ----------

func TestClearAuthCookie(t *testing.T) {
	rr := httptest.NewRecorder()

	ClearAuthCookie(rr)

	cookies := rr.Result().Cookies()

	if len(cookies) == 0 {
		t.Fatal("cookie not cleared")
	}

	if cookies[0].MaxAge != -1 {
		t.Fatal("cookie MaxAge not set to -1")
	}
}
