package main

import (
	"encoding/json"
	"github.com/archivers-space/archive"
	"net/http"
)

func UncrawlableHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		GetUncrawlableHandler(w, r)
	case "PUT":
		SaveUncrawlablesHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func UncrawlablesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ListUncrawlablesHandler(w, r)
	case "PUT", "POST":
		SaveUncrawlablesHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func GetUncrawlableHandler(w http.ResponseWriter, r *http.Request) {
	res := &archive.Uncrawlable{}
	args := &UncrawlablesGetArgs{
		Id:  r.URL.Path[len("/v0/uncrawlables/"):],
		Url: r.FormValue("url"),
	}
	err := new(Uncrawlables).Get(args, res)
	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, res)
}

func ListUncrawlablesHandler(w http.ResponseWriter, r *http.Request) {
	p := PageFromRequest(r)
	res := make([]*archive.Uncrawlable, p.Size)
	args := &UncrawlablesListArgs{
		Page:    p,
		OrderBy: "created",
	}
	err := new(Uncrawlables).List(args, &res)
	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	writePageResponse(w, res, r, p)
}

func SaveUncrawlablesHandler(w http.ResponseWriter, r *http.Request) {
	un := &archive.Uncrawlable{}
	if err := json.NewDecoder(r.Body).Decode(un); err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	res := &archive.Uncrawlable{}
	if err := new(Uncrawlables).Save(un, res); err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, res)
}
