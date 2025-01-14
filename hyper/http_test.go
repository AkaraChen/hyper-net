package hyper

import (
	"net/http"
	"net/url"
	"testing"
)

func TestHTTPHandlerIntoHyper(t *testing.T) {
	server := New(HyperOption{})
	server.Get(
		"/",
		HTTPHandlerIntoHyper(http.FileServer(http.Dir("."))),
	)

	tc := NewHyperTest(server)
	res := tc.Test(&http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Path: "/",
		},
	})
	if res.Code != 200 {
		t.Errorf("expected status code 200, got %d", res.Code)
	}
}
