package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	s := httptest.NewServer(NewServerRoutes())
	defer s.Close()

	res, err := http.Get(s.URL + "/healthcheck")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("status code mismatch. expected: %d, got: %d", http.StatusOK, res.StatusCode)
	}
}
