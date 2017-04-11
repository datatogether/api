package main

import (
	"github.com/archivers-space/archive"
	"net/rpc"
)

func init() {
	rpc.Register(new(Urls))
}

type Urls int

type UrlsGetArgs struct {
	Id   string
	Url  string
	Hash string
}

func (u *Urls) Get(args *UrlsGetArgs, res *archive.Url) (err error) {
	url := &archive.Url{
		Id:   args.Id,
		Url:  args.Url,
		Hash: args.Hash,
	}
	err = url.Read(appDB)
	if err != nil {
		return err
	}

	*res = *url
	return nil
}

type UrlsListArgs struct {
	OrderBy string
	Page
}

func (u *Urls) List(args *UrlsListArgs, res *[]*archive.Url) (err error) {
	urls, err := archive.ListUrls(appDB, args.Page.Size, args.Page.Offset())
	if err != nil {
		return err
	}
	*res = urls
	return nil
}
