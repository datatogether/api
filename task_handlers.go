package main

import (
	"github.com/datatogether/api/apiutil"
	"github.com/datatogether/task-mgmt/tasks"
	"net"
	"net/http"
	"net/rpc"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		GetTaskHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		ListTasksHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("tcp", cfg.TasksServiceUrl)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	cli := rpc.NewClient(conn)
	page := apiutil.PageFromRequest(r)
	p := &tasks.TasksListParams{
		Limit:  page.Size,
		Offset: page.Offset(),
	}
	reply := []*tasks.Task{}
	if err := cli.Call("TaskRequests.List", p, &reply); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, reply)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("tcp", cfg.TasksServiceUrl)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	cli := rpc.NewClient(conn)
	p := tasks.TasksGetParams{
		Id: r.FormValue("id"),
	}
	reply := &tasks.Task{}
	if err := cli.Call("TaskRequests.Get", p, reply); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, reply)
}
