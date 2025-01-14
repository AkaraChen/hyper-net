package hyper

import "net/http"

func HTTPHandlerIntoHyper(handler http.Handler) handlerFunc {
	fn := func(c *Context) {
		handler.ServeHTTP(c.Writer, c.Req)
	}
	return fn
}
