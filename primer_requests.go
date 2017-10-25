package main

import (
	"github.com/datatogether/core"
)

type Primers int

type PrimersGetArgs struct {
	Id string
}

func (u *Primers) Get(args *PrimersGetArgs, res *core.Primer) (err error) {
	p := &core.Primer{
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

func (u *Primers) List(args *PrimersListArgs, res *[]*core.Primer) (err error) {
	ps, err := core.ListPrimers(store, args.Limit, args.Offset)
	if err != nil {
		return err
	}
	*res = ps
	return nil
}
