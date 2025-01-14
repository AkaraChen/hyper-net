package hyper

import "net/http"

type Header struct {
	req *http.Request
	res http.ResponseWriter
}

func (h *Header) Get(name string) string {
	return h.req.Header.Get(name)
}

func (h *Header) Set(name string, value string) {
	h.res.Header().Set(name, value)
}
