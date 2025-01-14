package main

import (
	"github.com/akarachen/hyper-net/hyper"
)

func main() {
	server := hyper.New(hyper.HyperOption{})
	server.Get(
		"/health",
		func(c *hyper.Context) {
			c.JSON(map[string]interface{}{
				"status": "ok",
			})
		},
	)
	v1 := hyper.New(hyper.HyperOption{Group: "/v1"})
	v1.Get(
		"/hello",
		func(c *hyper.Context) {
			c.Text([]byte("Hello World!"))
		},
	)
	hyper.Start(":80")
}
