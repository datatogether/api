package main

import (
	"github.com/archivers-space/archive"
	"net/rpc"
)

func init() {
	rpc.Register(new(Urls))
}

type Urls int

type UrlsGetParams struct {
	Id   string
	Url  string
	Hash string
}

func (u *Urls) Get(p *UrlsGetParams, res *archive.Url) (err error) {
	url := &archive.Url{
		Id:   p.Id,
		Url:  p.Url,
		Hash: p.Hash,
	}
	err = url.Read(appDB)
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

func (u *Urls) List(p *UrlsListParams, res *[]*archive.Url) (err error) {
	urls, err := archive.ListUrls(appDB, p.Limit, p.Offset)
	if err != nil {
		return err
	}
	*res = urls
	return nil
}
