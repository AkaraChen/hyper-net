package hyper

import "net/http"

type IMatcher interface {
	Match(*http.Request) bool
}

type Matcher struct {
	Pattern string
	Method  []string
}

func (m Matcher) Match(r *http.Request) bool {
	if r.URL.Path != m.Pattern {
		return false
	}
	if len(m.Method) != 0 && !Contains(m.Method, r.Method) {
		return false
	}
	return true
}

type Route struct {
	Matcher IMatcher
	Handler func(*Context)
}
