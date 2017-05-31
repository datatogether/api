package main

import (
	"github.com/archivers-space/archive"
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
	args := &UrlsGetArgs{
		Id:   r.URL.Path[len("/urls/"):],
		Url:  r.FormValue("url"),
		Hash: r.FormValue("hash"),
	}
	err := new(Urls).Get(args, res)
	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, res)
}

func ListUrlsHandler(w http.ResponseWriter, r *http.Request) {
	p := PageFromRequest(r)
	res := make([]*archive.Url, p.Size)
	args := &UrlsListArgs{
		Page:    p,
		OrderBy: "created",
	}
	err := new(Urls).List(args, &res)
	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	writePageResponse(w, res, r, p)
}
