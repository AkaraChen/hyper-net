package hyper

import (
	"context"
	"fmt"
	"log"
)

type Logger struct {
	Middleware
}

func (l Logger) Handler(handler handlerFunc) handlerFunc {
	return func(c *Context) {
		log.Println(fmt.Println("Request from: ", c.Req.Host))
		handler(c)
		if err := recover(); err != nil {
			log.Fatalln(err)
			panic(err)
		}
	}
}

type EnvironmentContext struct {
	Middleware
}

func NewEnvironmentContextMiddleware(env string) MiddlewareFunc {
	return func(handler handlerFunc) handlerFunc {
		return func(c *Context) {
			c.Context = context.WithValue(c.Context, "environment", env)
			handler(c)
		}
	}
}
