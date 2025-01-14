package hyper

import (
	"fmt"
	"net/http"
)

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

func (h *Header) Del(name string) {
	h.res.Header().Del(name)
}

func (h *Header) Append(name string, value string) {
	v := h.res.Header().Get(name)
	if v == "" {
		v = value
	} else {
		v = fmt.Sprintf("%s,%s", v, value)
	}
	h.res.Header().Set(name, v)
}
