package hyper

import "net/http"

func (c *Context) Redirect(url string, permanent bool) {
	if permanent {
		c.Status(http.StatusMovedPermanently)
	} else {
		c.Status(http.StatusFound)
	}
	c.Writer.Header().Set("Location", url)
}
