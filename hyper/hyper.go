package hyper

import (
	"net/http"
)

type Hyper struct {
	routes []Route
}

func (h *Hyper) addRoute(matcher IMatcher, handler func(*Context)) {
	route := Route{
		Matcher: matcher,
		Handler: handler,
	}
	h.routes = append(h.routes, route)
}

func (h *Hyper) Handle(matcher string, handler func(*Context)) {
	h.addRoute(
		Matcher{
			Pattern: matcher,
			Method:  []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"},
		},
		handler,
	)
}

func (h *Hyper) Get(matcher string, handler func(*Context)) {
	h.addRoute(Matcher{Pattern: matcher, Method: []string{"GET"}}, handler)
}

func (h *Hyper) Post(matcher string, handler func(*Context)) {
	h.addRoute(Matcher{Pattern: matcher, Method: []string{"POST"}}, handler)
}

func (h *Hyper) Put(matcher string, handler func(*Context)) {
	h.addRoute(Matcher{Pattern: matcher, Method: []string{"PUT"}}, handler)
}

func (h *Hyper) Delete(matcher string, handler func(*Context)) {
	h.addRoute(Matcher{Pattern: matcher, Method: []string{"DELETE"}}, handler)
}

func (h *Hyper) Patch(matcher string, handler func(*Context)) {
	h.addRoute(Matcher{Pattern: matcher, Method: []string{"PATCH"}}, handler)
}

func (h *Hyper) Options(matcher string, handler func(*Context)) {
	h.addRoute(Matcher{Pattern: matcher, Method: []string{"OPTIONS"}}, handler)
}

func (h *Hyper) Head(matcher string, handler func(*Context)) {
	h.addRoute(Matcher{Pattern: matcher, Method: []string{"HEAD"}}, handler)
}

func (hyper *Hyper) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range hyper.routes {
		if route.Matcher.Match(req) {
			route.Handler(NewContext(w, req))
			break
		}
	}
}
