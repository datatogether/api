package main

import (
	"encoding/json"
	"github.com/datatogether/api/apiutil"
	"github.com/datatogether/archive"
	"net/http"
)

func CustomCrawlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		GetCustomCrawlHandler(w, r)
	case "PUT":
		SaveCustomCrawlHandler(w, r)
	case "DELETE":
		DeleteCustomCrawlHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func CustomCrawlsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ListCustomCrawlsHandler(w, r)
	case "PUT", "POST":
		SaveCustomCrawlHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func GetCustomCrawlHandler(w http.ResponseWriter, r *http.Request) {
	res := &archive.CustomCrawl{}
	args := &CustomCrawlsGetParams{
		Id: r.URL.Path[len("/customcrawls/"):],
	}
	err := new(CustomCrawls).Get(args, res)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, res)
}

func ListCustomCrawlsHandler(w http.ResponseWriter, r *http.Request) {
	p := apiutil.PageFromRequest(r)
	res := make([]*archive.CustomCrawl, p.Size)
	args := &CustomCrawlsListParams{
		Limit:   p.Limit(),
		Offset:  p.Offset(),
		OrderBy: "created",
	}
	err := new(CustomCrawls).List(args, &res)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WritePageResponse(w, res, r, p)
}

func SaveCustomCrawlHandler(w http.ResponseWriter, r *http.Request) {
	un := &archive.CustomCrawl{}
	if err := json.NewDecoder(r.Body).Decode(un); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	res := &archive.CustomCrawl{}
	if err := new(CustomCrawls).Save(un, res); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, res)
}

func DeleteCustomCrawlHandler(w http.ResponseWriter, r *http.Request) {
	un := &archive.CustomCrawl{}
	if err := json.NewDecoder(r.Body).Decode(un); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	res := &archive.CustomCrawl{}
	if err := new(CustomCrawls).Save(un, res); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, res)
}
