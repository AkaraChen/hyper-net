package hyper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeaderAppend(t *testing.T) {
	server := New(HyperOption{})
	server.Get("/append-header", func(c *Context) {
		c.Header.Append("X-Test-Header", "Value1")
		c.Header.Append("X-Test-Header", "Value2")
		c.Text([]byte("Headers appended"))
	})

	tc := NewHyperTest(server)
	req := httptest.NewRequest(http.MethodGet, "/append-header", nil)
	res := tc.Test(req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.Code)
	}

	headerValue := res.Header().Get("X-Test-Header")
	expectedValue := "Value1, Value2"
	if headerValue != expectedValue {
		t.Errorf("expected X-Test-Header value %q, got %q", expectedValue, headerValue)
	}

	if res.Body.String() != "Headers appended" {
		t.Errorf("expected body 'Headers appended', got %q", res.Body.String())
	}
}
