package main

import (
	"github.com/qri-io/archive"
	"net/http"
)

// ListPrimers
func ListPrimersHandler(w http.ResponseWriter, r *http.Request) {
	list, err := archive.ListPrimers(appDB, 25, 0)
	if err != nil {
		writeErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	// TODO - this should be writePageResponse
	writeResponse(w, list)
}
