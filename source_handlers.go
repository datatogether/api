package main

import (
	"github.com/qri-io/archive"
	"net/http"
)

func SourceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		GetSourceHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func SourcesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ListSourcesHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func GetSourceHandler(w http.ResponseWriter, r *http.Request) {
	res := &archive.Source{}
	args := &SourcesGetArgs{
		Id: r.URL.Path[len("/v0/sources/"):],
	}
	err := new(Sources).Get(args, res)
	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, res)
}

func ListSourcesHandler(w http.ResponseWriter, r *http.Request) {
	p := PageFromRequest(r)
	res := make([]*archive.Source, p.Size)
	args := &SourcesListArgs{
		Page:    p,
		OrderBy: "created",
	}
	err := new(Sources).List(args, &res)
	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	writePageResponse(w, res, r, p)
}
