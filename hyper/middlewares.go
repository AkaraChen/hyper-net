package hyper

type Middleware interface {
	Handler(handler handlerFunc) handlerFunc
}

type MiddlewareFunc func(handlerFunc) handlerFunc

func (m MiddlewareFunc) Handler(h handlerFunc) handlerFunc {
	return m(h)
}

// ApplyMiddlewares applies a list of middlewares to a handlerFunc, wrapping them in an onion pattern.
func ApplyMiddlewares(h handlerFunc, middlewares ...Middleware) handlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i].Handler(h)
	}
	return h
}
