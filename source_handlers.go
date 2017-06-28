package main

import (
	"github.com/datatogether/api/apiutil"
	"github.com/datatogether/archive"
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
	args := &SourcesGetParams{
		Id: r.URL.Path[len("/sources/"):],
	}
	err := new(Sources).Get(args, res)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, res)
}

func ListSourcesHandler(w http.ResponseWriter, r *http.Request) {
	p := apiutil.PageFromRequest(r)
	res := make([]*archive.Source, p.Size)
	args := &SourcesListParams{
		Limit:   p.Limit(),
		Offset:  p.Offset(),
		OrderBy: "created",
	}
	err := new(Sources).List(args, &res)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WritePageResponse(w, res, r, p)
}
