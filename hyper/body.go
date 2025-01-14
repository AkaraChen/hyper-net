package hyper

import (
	"encoding/json"
	"net/http"
)

type Body struct {
	req *http.Request
	res http.ResponseWriter
}

func (b *Body) Json(data interface{}) error {
	b.res.Header().Set(HeaderContentType, MIMEApplicationJSON)
	return json.NewEncoder(b.res).Encode(data)
}

func (b *Body) Text() (string, error) {
	var data []byte
	_, err := b.res.Write(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (b *Body) Form() {}
