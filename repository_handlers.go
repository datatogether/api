package main

import (
	"net/http"
)

func RepositoriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":

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
	default:
		NotFoundHandler(w, r)
	}
}
