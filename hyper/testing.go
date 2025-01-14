package hyper

import (
	"net/http"
	"net/http/httptest"
)

type HyperTest struct {
	*Hyper
}

func NewHyperTest(hyper *Hyper) *HyperTest {
	return &HyperTest{Hyper: hyper}
}

func (h *HyperTest) Test(req *http.Request) *httptest.ResponseRecorder {
	handler, _ := h.Mux.Handler(req)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	return res
}
