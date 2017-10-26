package main

import (
	"github.com/datatogether/core"
)

type Sources int

type SourcesGetParams struct {
	Id     string
	Source string
	Hash   string
}

func (u Sources) Get(args *SourcesGetParams, res *core.Source) (err error) {
	s := &core.Source{
		Id: args.Id,
	}
	err = s.Read(store)
	if err != nil {
		return err
	}

	err = s.Primer.Read(store)
	if err != nil {
		return err
	}

	*res = *s
	return nil
}

type SourcesListParams struct {
	OrderBy string
	Limit   int
	Offset  int
}

func (u Sources) List(p *SourcesListParams, res *[]*core.Source) (err error) {
	urls, err := core.ListSources(store, p.Limit, p.Offset)
	if err != nil {
		return err
	}
	*res = urls
	return nil
}
