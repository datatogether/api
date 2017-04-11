package main

import (
	"github.com/archivers-space/archive"
	"net/rpc"
)

func init() {
	rpc.Register(new(Primers))
}

type Primers int

type PrimersGetArgs struct {
	Id string
}

func (u *Primers) Get(args *PrimersGetArgs, res *archive.Primer) (err error) {
	p := &archive.Primer{
		Id: args.Id,
	}
	err = p.Read(appDB)
	if err != nil {
		return err
	}

	*res = *p
	return nil
}

type PrimersListArgs struct {
	OrderBy string
	Page
}

func (u *Primers) List(args *PrimersListArgs, res *[]*archive.Primer) (err error) {
	ps, err := archive.ListPrimers(appDB, args.Page.Size, args.Page.Offset())
	if err != nil {
		return err
	}
	*res = ps
	return nil
}
