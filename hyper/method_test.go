package hyper

import (
	"net/http"
	"net/url"
	"testing"
)

func TestSingleMethod(t *testing.T) {
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

func TestAllMethods(t *testing.T) {
	server := New(HyperOption{})
	server.All("/creation", func(c *Context) {
		c.Text([]byte("Hello World!"))
	})
	tc := NewHyperTest(server)
	for _, method := range methods {
		res := tc.Test(&http.Request{
			Method: method,
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
}
