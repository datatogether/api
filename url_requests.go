package main

import (
	"github.com/datatogether/core"
)

type Urls int

type UrlsGetParams struct {
	Id   string
	Url  string
	Hash string
}

func (u *Urls) Get(p *UrlsGetParams, res *core.Url) (err error) {
	url := &core.Url{
		Id:   p.Id,
		Url:  p.Url,
		Hash: p.Hash,
	}
	err = url.Read(store)
	if err != nil {
		return err
	}

	*res = *url
	return nil
}

type UrlsListParams struct {
	OrderBy string
	Limit   int
	Offset  int
}

func (u *Urls) List(p *UrlsListParams, res *[]*core.Url) (err error) {
	urls, err := core.ListUrls(store, p.Limit, p.Offset)
	if err != nil {
		return err
	}
	*res = urls
	return nil
}
