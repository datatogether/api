package main

import (
	"github.com/datatogether/api/apiutil"
	"github.com/datatogether/core"
	"github.com/datatogether/coverage/repositories"
	"net"
	"net/http"
	"net/rpc"
)

func RepositoriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		ListRepositoriesHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

// func GetRepositoryHandler(w http.ResponseWriter, r *http.Request) {
// 	// rpc.Dial("tcp", cfg.)
// }

func RepositoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		GetRepositoryHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func ListRepositoriesHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("tcp", cfg.CoverageServiceUrl)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	cli := rpc.NewClient(conn)
	p := repositories.RepositoryListParams{}
	reply := []*core.DataRepo{}
	if err := cli.Call("RepositoryRequests.List", p, &reply); err != nil {
		log.Info(err)
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, reply)
}

func GetRepositoryHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("tcp", cfg.CoverageServiceUrl)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	cli := rpc.NewClient(conn)
	p := repositories.RepositoryGetParams{
		Id: r.FormValue("id"),
	}
	reply := &core.DataRepo{}
	if err := cli.Call("RepositoryRequests.Get", p, &reply); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, reply)
}
