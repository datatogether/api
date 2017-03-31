package main

import (
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
	default:
		NotFoundHandler(w, r)
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		EmptyOkHandler(w, r)
	case "GET":
	default:
		NotFoundHandler(w, r)
	}
}
