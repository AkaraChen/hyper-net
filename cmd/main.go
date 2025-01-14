package main

import (
	"net/http"

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
	http.Handle("/", server)
	http.ListenAndServe(":80", nil)
}
