package hyper

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSetCookie(t *testing.T) {
	server := New(HyperOption{})
	server.Get("/set-cookie", func(c *Context) {
		options := &CookieOptions{
			Path:     "/",
			Domain:   "example.com",
			Expires:  time.Now().Add(24 * time.Hour),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
		c.SetCookie("test_cookie", "test_value", options)
		c.Text([]byte("Cookie set"))
	})

	tc := NewHyperTest(server)
	req := httptest.NewRequest(http.MethodGet, "/set-cookie", nil)
	res := tc.Test(req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.Code)
	}

	cookies := res.Result().Cookies()
	if len(cookies) != 1 {
		t.Fatalf("expected 1 cookie, got %d", len(cookies))
	}

	cookie := cookies[0]
	if cookie.Name != "test_cookie" {
		t.Errorf("expected cookie name 'test_cookie', got '%s'", cookie.Name)
	}
	if cookie.Value != "test_value" {
		t.Errorf("expected cookie value 'test_value', got '%s'", cookie.Value)
	}
	if cookie.Path != "/" {
		t.Errorf("expected cookie path '/', got '%s'", cookie.Path)
	}
	if cookie.Domain != "example.com" {
		t.Errorf("expected cookie domain 'example.com', got '%s'", cookie.Domain)
	}
	if !cookie.Secure {
		t.Error("expected cookie to be secure")
	}
	if !cookie.HttpOnly {
		t.Error("expected cookie to be HTTP only")
	}
	if cookie.SameSite != http.SameSiteStrictMode {
		t.Errorf("expected SameSite to be Strict, got %v", cookie.SameSite)
	}
}
