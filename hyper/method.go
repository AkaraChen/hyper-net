package hyper

import (
	"fmt"
	"net/http"
)

func (h *Hyper) handleMethod(method, path string, handler handlerFunc, middlewares ...Middleware) {
	pattern := fmt.Sprintf("%s %s%s", method, h.Group, path)
	h.Mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		defer recover()
		c := NewContext(w, r)
		m := append(h.middlewares, middlewares...)
		ApplyMiddlewares(handler, m...)(c)
	})
}

func (h *Hyper) Get(path string, handler handlerFunc, middlewares ...Middleware) {
	h.handleMethod(http.MethodGet, path, handler, middlewares...)
}

func (h *Hyper) Post(path string, handler handlerFunc, middlewares ...Middleware) {
	h.handleMethod(http.MethodPost, path, handler, middlewares...)
}

func (h *Hyper) Put(path string, handler handlerFunc, middlewares ...Middleware) {
	h.handleMethod(http.MethodPut, path, handler, middlewares...)
}

func (h *Hyper) Delete(path string, handler handlerFunc, middlewares ...Middleware) {
	h.handleMethod(http.MethodDelete, path, handler)
}

func (h *Hyper) Patch(path string, handler handlerFunc, middlewares ...Middleware) {
	h.handleMethod(http.MethodPatch, path, handler)
}

func (h *Hyper) Head(path string, handler handlerFunc, middlewares ...Middleware) {
	h.handleMethod(http.MethodHead, path, handler)
}

func (h *Hyper) Options(path string, handler handlerFunc, middlewares ...Middleware) {
	h.handleMethod(http.MethodOptions, path, handler)
}

func (h *Hyper) Trace(path string, handler handlerFunc, middlewares ...Middleware) {
	h.handleMethod(http.MethodTrace, path, handler)
}

func (h *Hyper) Connect(path string, handler handlerFunc, middlewares ...Middleware) {
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
