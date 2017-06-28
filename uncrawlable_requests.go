package main

import (
	"github.com/datatogether/archive"
)

type Uncrawlables int

type UncrawlablesGetParams struct {
	Id  string
	Url string
}

func (u *Uncrawlables) Get(p *UncrawlablesGetParams, res *archive.Uncrawlable) (err error) {
	url := &archive.Uncrawlable{
		Id:  p.Id,
		Url: p.Url,
	}
	err = url.Read(store)
	if err != nil {
		return err
	}

	*res = *url
	return nil
}

type UncrawlablesListParams struct {
	OrderBy string
	Limit   int
	Offset  int
}

func (u *Uncrawlables) List(p *UncrawlablesListParams, res *[]*archive.Uncrawlable) (err error) {
	urls, err := archive.ListUncrawlables(store, p.Limit, p.Offset)
	if err != nil {
		return err
	}
	*res = urls
	return nil
}

func (u *Uncrawlables) Save(model *archive.Uncrawlable, res *archive.Uncrawlable) (err error) {
	err = model.Save(store)
	if err != nil {
		return err
	}

	*res = *model
	return nil
}

func (u *Uncrawlables) Delete(model *archive.Uncrawlable, res *archive.Uncrawlable) (err error) {
	err = model.Delete(store)
	if err != nil {
		return err
	}

	*res = *model
	return nil
}
