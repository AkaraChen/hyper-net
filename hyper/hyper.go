package hyper

import (
	"net/http"
	"runtime"
	"strings"
)

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
}

func (h *Hyper) Start(addr string) error {
	return http.ListenAndServe(addr, nil)
}
