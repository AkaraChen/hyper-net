package hyper

import (
	"fmt"
	"net/http"
)

func (h *Hyper) handleMethod(method, path string, handler handlerFunc) {
	pattern := fmt.Sprintf("%s %s%s", method, h.Group, path)
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		c := &Context{Writer: w, Req: r}
		handler(c)
	})
}

func (h *Hyper) Get(path string, handler handlerFunc) {
	h.handleMethod(http.MethodGet, path, handler)
}

func (h *Hyper) Post(path string, handler handlerFunc) {
	h.handleMethod(http.MethodPost, path, handler)
}

func (h *Hyper) Put(path string, handler handlerFunc) {
	h.handleMethod(http.MethodPut, path, handler)
}

func (h *Hyper) Delete(path string, handler handlerFunc) {
	h.handleMethod(http.MethodDelete, path, handler)
}

func (h *Hyper) Patch(path string, handler handlerFunc) {
	h.handleMethod(http.MethodPatch, path, handler)
}

func (h *Hyper) Head(path string, handler handlerFunc) {
	h.handleMethod(http.MethodHead, path, handler)
}

func (h *Hyper) Options(path string, handler handlerFunc) {
	h.handleMethod(http.MethodOptions, path, handler)
}

func (h *Hyper) Trace(path string, handler handlerFunc) {
	h.handleMethod(http.MethodTrace, path, handler)
}

func (h *Hyper) Connect(path string, handler handlerFunc) {
	h.handleMethod(http.MethodConnect, path, handler)
}

var methods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
	http.MethodPatch,
	http.MethodHead,
	http.MethodOptions,
	http.MethodTrace,
	http.MethodConnect,
}

func (h *Hyper) All(path string, handler handlerFunc) {
	for _, method := range methods {
		h.handleMethod(method, path, handler)
	}
}
