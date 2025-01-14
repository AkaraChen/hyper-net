package hyper

import "fmt"

func (c *Context) Download(filename string) {
	c.Header.Set(HeaderContentDisposition, "attachment")
	c.Header.Append(HeaderContentDisposition, fmt.Sprintf("filename=%s", filename))
}
