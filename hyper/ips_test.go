package hyper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestIP(t *testing.T) {
	server := New(HyperOption{})
	server.Get("/ip", func(c *Context) {
		ips := c.IP()
		c.JSON(ips)
	})

	tc := NewHyperTest(server)

	// Test with X-Forwarded-For header
	req := httptest.NewRequest(http.MethodGet, "/ip", nil)
	req.Header.Set("X-Forwarded-For", "192.168.1.1, 10.0.0.1")
	req.RemoteAddr = "127.0.0.1:12345"

	res := tc.Test(req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.Code)
	}

	var ips []string
	err := json.Unmarshal(res.Body.Bytes(), &ips)
	if err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}

	expectedIPs := []string{"127.0.0.1", "192.168.1.1", "10.0.0.1"}
	if !reflect.DeepEqual(ips, expectedIPs) {
		t.Errorf("expected IPs %v, got %v", expectedIPs, ips)
	}

	// Test without X-Forwarded-For header
	req = httptest.NewRequest(http.MethodGet, "/ip", nil)
	req.RemoteAddr = "192.168.0.1:54321"

	res = tc.Test(req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.Code)
	}

	err = json.Unmarshal(res.Body.Bytes(), &ips)
	if err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}

	expectedIPs = []string{"192.168.0.1"}
	if !reflect.DeepEqual(ips, expectedIPs) {
		t.Errorf("expected IPs %v, got %v", expectedIPs, ips)
	}
}
