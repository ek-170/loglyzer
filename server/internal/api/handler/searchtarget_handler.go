package handler

import (
	"fmt"
	"io"
	"net/http"
)

type SearchTargetHandler interface {
	HandleSearchTarget(w http.ResponseWriter, r *http.Request)
}

type searchTargetHandler struct {
}

func NewSearchTargetHandler() SearchTargetHandler {
	return &searchTargetHandler{}
}

func (ssh *searchTargetHandler) HandleSearchTarget(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
    case http.MethodGet:
        getSearchTarget(w, r)
    // case http.MethodPost:
    //     findSearchtarget(w, r)
    // case http.MethodPut:
    //     upsertSearchtarget(w, r)
    // case http.MethodDelete:
    //     deleteSearchtarget(w, r)
	default:
		http.Error(w, fmt.Sprintf("Method %s is not allowed", r.Method), 405)
	}
}

func getSearchTarget(w http.ResponseWriter, _ *http.Request) {
  io.WriteString(w, "Searchtarget World!\n")
}