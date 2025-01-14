package main

import (
	"fmt"

	"github.com/akarachen/hyper-net/hyper"
)

func main() {
	server := hyper.New(hyper.HyperOption{
		Middlewares: []hyper.Middleware{
			hyper.Logger{},
		},
	})
	server.Get(
		"/",
		func(c *hyper.Context) {
			c.Header.Set(hyper.HeaderContentType, hyper.MIMETextPlain)
			c.Text([]byte("<h1>Hello World!</h1>"))
		},
	)
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
		"/hello/{name}",
		func(c *hyper.Context) {
			c.Text([]byte(fmt.Sprintf("Hello %s!", c.PathValue("name"))))
		},
	)
	server.Mount(v1)
	server.Start(":80")
}
