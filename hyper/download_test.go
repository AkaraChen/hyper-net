package hyper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDownload(t *testing.T) {
	server := New(HyperOption{})
	server.Get("/download", func(c *Context) {
		c.Download("test.txt")
		c.Text([]byte("File content"))
	})

	tc := NewHyperTest(server)
	req := httptest.NewRequest(http.MethodGet, "/download", nil)
	res := tc.Test(req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.Code)
	}

	contentDisposition := res.Header().Get("Content-Disposition")
	expectedDisposition := `attachment; filename=test.txt`
	if contentDisposition != expectedDisposition {
		t.Errorf("expected Content-Disposition header %q, got %q", expectedDisposition, contentDisposition)
	}

	if res.Body.String() != "File content" {
		t.Errorf("expected body 'File content', got %q", res.Body.String())
	}
}
