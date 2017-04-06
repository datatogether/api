package main

import (
	"github.com/qri-io/archive"
	"net/rpc"
)

func init() {
	rpc.Register(new(Uncrawlables))
}

type Uncrawlables int

type UncrawlablesGetArgs struct {
	Id  string
	Url string
}

func (u *Uncrawlables) Get(args *UncrawlablesGetArgs, res *archive.Uncrawlable) (err error) {
	url := &archive.Uncrawlable{
		Id:  args.Id,
		Url: args.Url,
	}
	err = url.Read(appDB)
	if err != nil {
		return err
	}

	*res = *url
	return nil
}

type UncrawlablesListArgs struct {
	OrderBy string
	Page
}

func (u *Uncrawlables) List(args *UncrawlablesListArgs, res *[]*archive.Uncrawlable) (err error) {
	urls, err := archive.ListUncrawlables(appDB, args.Page.Size, args.Page.Offset())
	if err != nil {
		return err
	}
	*res = urls
	return nil
}
