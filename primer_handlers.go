package main

import (
	"github.com/archivers-space/archive"
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
		Id: r.URL.Path[len("/v0/primers/"):],
	}
	err := new(Primers).Get(args, res)
	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, res)
}

func ListPrimersHandler(w http.ResponseWriter, r *http.Request) {
	p := PageFromRequest(r)
	res := make([]*archive.Primer, p.Size)
	args := &PrimersListArgs{
		Page:    p,
		OrderBy: "created",
	}
	err := new(Primers).List(args, &res)
	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	writePageResponse(w, res, r, p)
}
