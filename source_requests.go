package main

import (
	"github.com/qri-io/archive"
	"net/rpc"
)

func init() {
	rpc.Register(new(Sources))
}

type Sources int

type SourcesGetArgs struct {
	Id     string
	Source string
	Hash   string
}

func (u *Sources) Get(args *SourcesGetArgs, res *archive.Source) (err error) {
	s := &archive.Source{
		Id: args.Id,
	}
	err = s.Read(appDB)
	if err != nil {
		return err
	}

	err = s.Primer.Read(appDB)
	if err != nil {
		return err
	}

	*res = *s
	return nil
}

type SourcesListArgs struct {
	OrderBy string
	Page
}

func (u *Sources) List(args *SourcesListArgs, res *[]*archive.Source) (err error) {
	urls, err := archive.ListSources(appDB, args.Page.Size, args.Page.Offset())
	if err != nil {
		return err
	}
	*res = urls
	return nil
}
