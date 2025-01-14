package hyper

import (
	"fmt"
	"net/http"
)

type handlerFunc func(*Context)

func (h *Hyper) handleMethod(method, path string, handler handlerFunc) {
	pattern := fmt.Sprintf("%s %s", method, path)
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		c := &Context{Writer: w, Req: r}
		handler(c)
	})
}

func (h *Hyper) Get(path string, handler handlerFunc) {
	h.handleMethod("GET", path, handler)
}

func (h *Hyper) Post(path string, handler handlerFunc) {
	h.handleMethod("POST", path, handler)
}

func (h *Hyper) Put(path string, handler handlerFunc) {
	h.handleMethod("PUT", path, handler)
}

func (h *Hyper) Delete(path string, handler handlerFunc) {
	h.handleMethod("DELETE", path, handler)
}

func (h *Hyper) Patch(path string, handler handlerFunc) {
	h.handleMethod("PATCH", path, handler)
}

func (h *Hyper) Head(path string, handler handlerFunc) {
	h.handleMethod("HEAD", path, handler)
}

func (h *Hyper) Options(path string, handler handlerFunc) {
	h.handleMethod("OPTIONS", path, handler)
}
