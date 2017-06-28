package main

import (
	"encoding/json"
	"github.com/datatogether/api/apiutil"
	"github.com/datatogether/archive"
	"net/http"
)

func UncrawlableHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		GetUncrawlableHandler(w, r)
	case "PUT":
		SaveUncrawlableHandler(w, r)
	case "DELETE":
		DeleteUncrawlableHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func UncrawlablesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ListUncrawlablesHandler(w, r)
	case "PUT", "POST":
		SaveUncrawlableHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func GetUncrawlableHandler(w http.ResponseWriter, r *http.Request) {
	res := &archive.Uncrawlable{}
	args := &UncrawlablesGetParams{
		Id:  r.URL.Path[len("/uncrawlables/"):],
		Url: r.FormValue("url"),
	}
	err := new(Uncrawlables).Get(args, res)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, res)
}

func ListUncrawlablesHandler(w http.ResponseWriter, r *http.Request) {
	p := apiutil.PageFromRequest(r)
	res := make([]*archive.Uncrawlable, p.Size)
	args := &UncrawlablesListParams{
		Limit:   p.Limit(),
		Offset:  p.Offset(),
		OrderBy: "created",
	}
	err := new(Uncrawlables).List(args, &res)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WritePageResponse(w, res, r, p)
}

func SaveUncrawlableHandler(w http.ResponseWriter, r *http.Request) {
	un := &archive.Uncrawlable{}
	if err := json.NewDecoder(r.Body).Decode(un); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	res := &archive.Uncrawlable{}
	if err := new(Uncrawlables).Save(un, res); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, res)
}

func DeleteUncrawlableHandler(w http.ResponseWriter, r *http.Request) {
	un := &archive.Uncrawlable{}
	if err := json.NewDecoder(r.Body).Decode(un); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	res := &archive.Uncrawlable{}
	if err := new(Uncrawlables).Save(un, res); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, res)
}
