package hyper

import (
	"net/http"
	"net/url"
	"testing"
)

func TestNew(t *testing.T) {
	s := New(HyperOption{})
	if s == nil {
		t.Errorf("expected hyper, got nil")
	}
	if s.Group != "" {
		t.Errorf("expected empty group, got %s", s.Group)
	}
	s.Get("/", func(c *Context) {
		c.Text([]byte("Hello World!"))
	})
	v1 := New(HyperOption{Group: "/v1"})
	v1.Get("/", func(c *Context) {
		c.Text([]byte("Hello sub group!"))
	})
	s.Mount(v1)

	tc := NewHyperTest(s)

	res := tc.Test(&http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Path: "/",
		},
	})
	if res.Code != 200 {
		t.Errorf("expected status code 200, got %d", res.Code)
	}
	if res.Body.String() != "Hello World!" {
		t.Errorf("expected body Hello World!, got %s", res.Body.String())
	}

	res = tc.Test(&http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Path: "/v1/",
		},
	})
	if res.Code != 200 {
		t.Errorf("expected status code 200, got %d", res.Code)
	}
	if res.Body.String() != "Hello sub group!" {
		t.Errorf("expected body Hello sub group!, got %s", res.Body.String())
	}
}
