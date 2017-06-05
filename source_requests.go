package main

import (
	"github.com/archivers-space/archive"
)

type Sources int

type SourcesGetParams struct {
	Id     string
	Source string
	Hash   string
}

func (u Sources) Get(args *SourcesGetParams, res *archive.Source) (err error) {
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

type SourcesListParams struct {
	OrderBy string
	Limit   int
	Offset  int
}

func (u Sources) List(p *SourcesListParams, res *[]*archive.Source) (err error) {
	urls, err := archive.ListSources(appDB, p.Limit, p.Offset)
	if err != nil {
		return err
	}
	*res = urls
	return nil
}
