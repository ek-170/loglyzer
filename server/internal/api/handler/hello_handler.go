package handler

import (
	"fmt"
	"io"
	"net/http"
)

type HelloHandler interface {
	HandleHello(w http.ResponseWriter, r *http.Request)
}

type helloHandler struct {
}

func NewHelloHandler() HelloHandler {
	return &helloHandler{}
}

func (dh *helloHandler) HandleHello(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		hello(w, r)
	default:
		http.Error(w, fmt.Sprintf("Method %s is not allowed", r.Method), 405)
	}
}

func hello(w http.ResponseWriter, _ *http.Request) {
  io.WriteString(w, "Hello World!\n")
}