package main

import (
	"github.com/archivers-space/archive"
)

type Collections int

type CollectionsGetParams struct {
	Id string
}

func (u *Collections) Get(args *CollectionsGetParams, res *archive.Collection) (err error) {
	p := &archive.Collection{
		Id: args.Id,
	}
	err = p.Read(store)
	if err != nil {
		return err
	}

	*res = *p
	return nil
}

type CollectionsListParams struct {
	OrderBy string
	Limit   int
	Offset  int
}

func (u *Collections) List(args *CollectionsListParams, res *[]*archive.Collection) (err error) {
	ps, err := archive.ListCollections(store, args.Limit, args.Offset)
	if err != nil {
		return err
	}
	*res = ps
	return nil
}
