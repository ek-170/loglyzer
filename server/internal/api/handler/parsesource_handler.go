package handler

import (
	"fmt"
	"io"
	"net/http"
)

type ParseSourceHandler interface {
	HandleParseSource(w http.ResponseWriter, r *http.Request)
}

type parseSourceHandler struct {
}

func NewParseSourceHandler() ParseSourceHandler {
	return &parseSourceHandler{}
}

func (lh *parseSourceHandler) HandleParseSource(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
    case http.MethodGet:
        getParseSource(w, r)
    // case http.MethodPost:
    //     findParseSource(w, r)
    // case http.MethodPut:
    //     upsertParseSource(w, r)
    // case http.MethodDelete:
    //     deleteParseSource(w, r)
	default:
		http.Error(w, fmt.Sprintf("Method %s is not allowed", r.Method), 405)
	}
}

func getParseSource(w http.ResponseWriter, _ *http.Request) {
  io.WriteString(w, "ParseSource World!\n")
}