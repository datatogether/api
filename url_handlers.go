package main

import (
	"github.com/datatogether/api/apiutil"
	"github.com/datatogether/archive"
	"net/http"
)

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		GetUrlHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func UrlsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ListUrlsHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func GetUrlHandler(w http.ResponseWriter, r *http.Request) {
	res := &archive.Url{}
	args := &UrlsGetParams{
		Id:   r.URL.Path[len("/urls/"):],
		Url:  r.FormValue("url"),
		Hash: r.FormValue("hash"),
	}
	err := new(Urls).Get(args, res)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, res)
}

func ListUrlsHandler(w http.ResponseWriter, r *http.Request) {
	p := apiutil.PageFromRequest(r)
	res := make([]*archive.Url, p.Size)
	args := &UrlsListParams{
		Limit:   p.Limit(),
		Offset:  p.Offset(),
		OrderBy: "created",
	}
	err := new(Urls).List(args, &res)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WritePageResponse(w, res, r, p)
}
