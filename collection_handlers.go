package main

import (
	"github.com/archivers-space/api/apiutil"
	"github.com/archivers-space/archive"
	"net/http"
)

func CollectionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		GetCollectionHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func CollectionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ListCollectionsHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func GetCollectionHandler(w http.ResponseWriter, r *http.Request) {
	res := &archive.Collection{}
	args := &CollectionsGetParams{
		Id: r.URL.Path[len("/collections/"):],
		// Collection: r.FormValue("collection"),
		// Hash:       r.FormValue("hash"),
	}
	err := new(Collections).Get(args, res)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, res)
}

func ListCollectionsHandler(w http.ResponseWriter, r *http.Request) {
	p := apiutil.PageFromRequest(r)
	res := make([]*archive.Collection, p.Size)
	args := &CollectionsListParams{
		Limit:   p.Limit(),
		Offset:  p.Offset(),
		OrderBy: "created",
	}
	err := new(Collections).List(args, &res)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WritePageResponse(w, res, r, p)
}
