package hyper

import (
	"net/http"
	"time"
)

type CookieOptions struct {
	Path     string
	Domain   string
	Expires  time.Time
	Secure   bool
	HttpOnly bool
	SameSite http.SameSite
}

func (c *Context) SetCookie(name, value string, options *CookieOptions) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     options.Path,
		Domain:   options.Domain,
		Expires:  options.Expires,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
		SameSite: options.SameSite,
	})
}
