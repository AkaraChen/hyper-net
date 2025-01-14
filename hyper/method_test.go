package hyper

import (
	"net/http"
	"net/url"
	"testing"
)

func TestCreation(t *testing.T) {
	server := New(HyperOption{})
	server.Get("/creation", func(c *Context) {
		c.Text([]byte("Hello World!"))
	})
	tc := NewHyperTest(server)
	res := tc.Test(&http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Path: "/creation",
		},
	})
	if res.Code != 200 {
		t.Errorf("expected status code 200, got %d", res.Code)
	}
	if res.Body.String() != "Hello World!" {
		t.Errorf("expected body Hello World!, got %s", res.Body.String())
	}
}
