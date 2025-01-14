package hyper

import (
	"testing"
)

func TestCreation(t *testing.T) {
	server := new(Hyper)
	server.Get("/creation", func(c *Context) {
		c.Text([]byte("Hello World!"))
	})
}
