package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

// Proxied User model. The real user model is in github.com/archivers-space/identity/user.go
type User struct {
	Id          string `json:"id" sql:"id"`
	Created     int64  `json:"created" sql:"created"`
	Updated     int64  `json:"updated" sql:"updated"`
	Username    string `json:"username" sql:"username"`
	Email       string `json:"email" sql:"email"`
	Name        string `json:"name" sql:"name"`
	Description string `json:"description" sql:"description"`
	HomeUrl     string `json:"home_url" sql:"home_url"`
	CurrentKey  string `json:"currentKey"`
	Anonymous   bool   `json:"-"`
}

func requestAddUser(r *http.Request) (*http.Request, error) {
	u := anonymousUser(r)

	token := r.FormValue("api_token")
	if token != "" {
		res, err := http.Get(fmt.Sprintf("%s/users/?access_token=%s&envelope=false", cfg.IdentityServerUrl, token))
		if err != nil {
			logger.Println(err.Error())
			return r, err
		}
		if res.StatusCode == http.StatusOK {
			authUser := &User{}
			if err := json.NewDecoder(res.Body).Decode(authUser); err != nil {
				logger.Println(err.Error())
				return r, err
			}
			u = authUser
		}
	}

	ctx := r.Context()
	if u != nil {
		ctx = context.WithValue(ctx, "user", u)
	}

	return r.WithContext(ctx), nil
}

func getIP(r *http.Request) string {
	remoteAddr := r.Header.Get("x-forwarded-for")
	if remoteAddr != "" {
		return strings.TrimSpace(strings.Split(remoteAddr, ",")[0])
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}

	return ip
}

func anonymousUser(r *http.Request) *User {
	return &User{
		Username:  getIP(r),
		Anonymous: true,
	}
}
