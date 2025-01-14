package hyper

import "net/http"

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
	}
}

func (c *Context) Text(data []byte) (int, error) {
	return c.Writer.Write(data)
}

type Handler interface {
	ServeHTTP(*Context)
}
