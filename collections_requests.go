package main

import (
	"github.com/datatogether/core"
)

type Collections int

type CollectionsGetParams struct {
	Id string
}

func (u *Collections) Get(args *CollectionsGetParams, res *core.Collection) (err error) {
	p := &core.Collection{
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

func (u *Collections) List(args *CollectionsListParams, res *[]*core.Collection) (err error) {
	ps, err := core.ListCollections(store, args.Limit, args.Offset)
	if err != nil {
		return err
	}
	*res = ps
	return nil
}
