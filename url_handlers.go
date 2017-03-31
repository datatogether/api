package main

import (
	"github.com/qri-io/archive"
	"net/http"
	"strings"
)

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		logger.Println(r.Context().Value("user"))
		url := r.FormValue("url")
		u := &archive.Url{Url: url, Id: strings.TrimPrefix(r.URL.Path, "/urls/")}
		if err := u.Read(appDB); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Println(err.Error())
			return
		}

		writeResponse(w, url)
	default:
		NotFoundHandler(w, r)
	}
}

func UrlsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// if we have a "url" param, read that single url
		url := r.FormValue("url")
		if url != "" {
			u := &archive.Url{Url: url}
			if err := u.Read(appDB); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				logger.Println(err.Error())
				return
			}

			writeResponse(w, u)
		} else {
			p := PageFromRequest(r)
			var (
				urls []*archive.Url
				err  error
			)
			if fetched, _ := reqParamBool("fetched", r); fetched {
				urls, err = archive.FetchedUrls(appDB, p.Size, p.Offset())
			} else {
				urls, err = archive.ListUrls(appDB, p.Size, p.Offset())
			}
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				logger.Println(err.Error())
				return
			}

			writeResponse(w, urls)
		}
	default:
		NotFoundHandler(w, r)
	}
}
