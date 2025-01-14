package hyper

func (c *Context) Status(code int) *Context {
	c.Writer.WriteHeader(code)
	return c
}
