package hyper

import "fmt"

func (c *Context) Download(filename string) {
	c.Header.Set(HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", filename))
}
