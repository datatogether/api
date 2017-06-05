package main

import (
	"github.com/archivers-space/api/apiutil"
	"github.com/archivers-space/identity/user"
	"net"
	"net/http"
	"net/rpc"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		GetUserHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
		ListUsersHandler(w, r)
	default:
		NotFoundHandler(w, r)
	}
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("tcp", cfg.IdentityServiceUrl)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	cli := rpc.NewClient(conn)
	page := apiutil.PageFromRequest(r)
	p := user.UsersListParams{
		Limit:  page.Size,
		Offset: page.Offset(),
	}
	reply := []*user.User{}
	if err := cli.Call("UserRequests.List", p, &reply); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, reply)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("tcp", cfg.IdentityServiceUrl)
	if err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	cli := rpc.NewClient(conn)
	p := user.UsersGetParams{
		Subject: &user.User{
			Id: r.FormValue("id"),
		},
	}
	reply := &user.User{}
	if err := cli.Call("UserRequests.Get", p, reply); err != nil {
		apiutil.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	apiutil.WriteResponse(w, reply)
}
