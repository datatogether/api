package main

import (
	"github.com/datatogether/api/apiutil"
	"github.com/datatogether/archive"
	"net/http"
)

func PrimerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		GetPrimerHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func PrimersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ListPrimersHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func GetPrimerHandler(w http.ResponseWriter, r *http.Request) {
	res := &archive.Primer{}
	args := &PrimersGetArgs{
		Id: r.URL.Path[len("/primers/"):],
	}
	err := new(Primers).Get(args, res)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, res)
}

func ListPrimersHandler(w http.ResponseWriter, r *http.Request) {
	p := apiutil.PageFromRequest(r)
	res := make([]*archive.Primer, p.Size)
	args := &PrimersListArgs{
		Limit:   p.Limit(),
		Offset:  p.Offset(),
		OrderBy: "created",
	}
	err := new(Primers).List(args, &res)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WritePageResponse(w, res, r, p)
}
