package hyper

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	Header *Header
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Header: &Header{req: r, res: w},
	}
}

func (c *Context) Text(data []byte) (int, error) {
	return c.Writer.Write(data)
}

func (c *Context) JSON(data interface{}) error {
	c.Header.Set(HeaderContentType, MIMEApplicationJSON)
	return json.NewEncoder(c.Writer).Encode(data)
}

func (c *Context) PathValue(name string) string {
	return c.Req.PathValue(name)
}
