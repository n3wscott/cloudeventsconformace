package http

import (
	"net/http"
)

type Request struct {
	Header http.Header
	Body   []byte
}
