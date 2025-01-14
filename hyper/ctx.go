package hyper

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
)

type Context struct {
	Writer    http.ResponseWriter
	Req       *http.Request
	Header    *Header
	Body      *Body
	Method    string
	Query     url.Values
	Cookie    []*http.Cookie
	FormValue func(key string) string
	FormFile  func(key string) (multipart.File, *multipart.FileHeader, error)
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:    w,
		Req:       r,
		Header:    &Header{req: r, res: w},
		Body:      &Body{req: r, res: w},
		Method:    r.Method,
		Query:     r.URL.Query(),
		Cookie:    r.Cookies(),
		FormValue: r.FormValue,
		FormFile:  r.FormFile,
	}
}

func (c *Context) Text(data []byte) (int, error) {
	return c.Writer.Write(data)
}

func (c *Context) JSON(data interface{}) error {
	c.Header.Set(HeaderContentType, MIMEApplicationJSON)
	return json.NewEncoder(c.Writer).Encode(data)
}

func (c *Context) PathValue(name string) string {
	return c.Req.PathValue(name)
}
