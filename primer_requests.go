package main

import (
	"github.com/datatogether/archive"
)

type Primers int

type PrimersGetArgs struct {
	Id string
}

func (u *Primers) Get(args *PrimersGetArgs, res *archive.Primer) (err error) {
	p := &archive.Primer{
		Id: args.Id,
	}
	err = p.Read(store)
	if err != nil {
		return err
	}

	*res = *p
	return nil
}

type PrimersListArgs struct {
	OrderBy string
	Limit   int
	Offset  int
}

func (u *Primers) List(args *PrimersListArgs, res *[]*archive.Primer) (err error) {
	ps, err := archive.ListPrimers(store, args.Limit, args.Offset)
	if err != nil {
		return err
	}
	*res = ps
	return nil
}
