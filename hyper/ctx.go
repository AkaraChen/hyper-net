package hyper

import "net/http"

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
}

func (c *Context) Text(data []byte) (int, error) {
	return c.Writer.Write(data)
}
