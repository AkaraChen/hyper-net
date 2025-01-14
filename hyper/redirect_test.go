package hyper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirect(t *testing.T) {
	server := New(HyperOption{})
	server.Get("/redirect", func(c *Context) {
		c.Redirect("/target", false)
	})

	tc := NewHyperTest(server)
	req := httptest.NewRequest(http.MethodGet, "/redirect", nil)
	res := tc.Test(req)

	if res.Code != http.StatusFound {
		t.Errorf("expected status code %d, got %d", http.StatusFound, res.Code)
	}

	location := res.Header().Get("Location")
	if location != "/target" {
		t.Errorf("expected Location header to be '/target', got %q", location)
	}

	// Test permanent redirect
	server.Get("/permanent-redirect", func(c *Context) {
		c.Redirect("/permanent-target", true)
	})

	req = httptest.NewRequest(http.MethodGet, "/permanent-redirect", nil)
	res = tc.Test(req)

	if res.Code != http.StatusMovedPermanently {
		t.Errorf("expected status code %d, got %d", http.StatusMovedPermanently, res.Code)
	}

	location = res.Header().Get("Location")
	if location != "/permanent-target" {
		t.Errorf("expected Location header to be '/permanent-target', got %q", location)
	}
}
