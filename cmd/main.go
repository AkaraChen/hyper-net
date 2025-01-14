package main

import (
	"github.com/akarachen/hyper-net/hyper"
)

func main() {
	server := new(hyper.Hyper)
	server.Get(
		"/",
		func(c *hyper.Context) {
			c.Text([]byte("Hello World!"))
		},
	)
	server.Get(
		"/health",
		func(c *hyper.Context) {
			c.Text([]byte("OK"))
		},
	)
	server.Get(
		"/greet/{name}",
		func(c *hyper.Context) {
			c.Text([]byte("Hello " + c.Req.PathValue("name")))
		},
	)
	server.Start(":80")
}
