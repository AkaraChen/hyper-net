package hyper

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
}

func (c *Context) Text(data []byte) (int, error) {
	return c.Writer.Write(data)
}

func (c *Context) JSON(data interface{}) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(c.Writer).Encode(data)
}

func (c *Context) PathValue(name string) string {
	return c.Req.PathValue(name)
}
