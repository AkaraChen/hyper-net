package hyper

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/ory/graceful"
)

type handlerFunc func(*Context)

func init() {
	v := runtime.Version()
	semver := strings.Split(v, ".")
	major := semver[0]
	minor := semver[1]
	// should upper than 1.22
	if major < "1" || major == "1" && minor < "22" {
		panic("hyper-net only support golang version >= 1.22")
	}
}

type Hyper struct {
	Group       string
	Mux         *http.ServeMux
	middlewares []Middleware
}

type HyperOption struct {
	Group       string
	Middlewares []Middleware
}

func New(opts HyperOption) *Hyper {
	h := new(Hyper)
	if opts.Group != "" {
		h.Group = opts.Group
	}
	h.Mux = http.NewServeMux()
	h.middlewares = opts.Middlewares
	return h
}

func (h *Hyper) Mount(hyper *Hyper) {
	pattern := fmt.Sprintf("%s%s/", h.Group, hyper.Group)
	for _, method := range methods {
		h.Mux.Handle(fmt.Sprintf("%s %s", method, pattern), hyper.Mux)
	}
}

func (h *Hyper) Start(addr string) {
	server := graceful.WithDefaults(&http.Server{
		Addr:    addr,
		Handler: h.Mux,
	})
	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		log.Fatalln("main: Failed to gracefully shutdown")
	}
	log.Println("main: Server was shutdown gracefully")
}
