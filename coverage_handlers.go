package main

import (
	"github.com/archivers-space/api/apiutil"
	"github.com/archivers-space/coverage/coverage"
	"net"
	"net/http"
	"net/rpc"
)

func CoverageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		GetCoverageHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

// func CoveragesHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "OPTIONS":
// 		EmptyOkHandler(w, r)
// 	case "GET":
// 		ListCoveragesHandler(w, r)
// 	default:
// 		NotFoundHandler(w, r)
// 	}
// }

// func ListCoveragesHandler(w http.ResponseWriter, r *http.Request) {
// 	conn, err := net.Dial("tcp", cfg.CoverageServiceUrl)
// 	if err != nil {
// 		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	cli := rpc.NewClient(conn)
// 	page := apiutil.PageFromRequest(r)
// 	p := user.CoveragesListParams{
// 		Limit:  page.Size,
// 		Offset: page.Offset(),
// 	}
// 	reply := []*user.Coverage{}
// 	if err := cli.Call("CoverageRequests.List", p, &reply); err != nil {
// 		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	apiutil.WriteResponse(w, reply)
// }

func GetCoverageHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("tcp", cfg.CoverageServiceUrl)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	cli := rpc.NewClient(conn)
	p := coverage.CoverageSummaryParams{
		Pattern: r.FormValue("pattern"),
	}
	reply := &coverage.Summary{}
	if err := cli.Call("CoverageRequests.Summary", p, reply); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, reply)
}
