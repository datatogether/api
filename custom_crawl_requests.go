package main

import (
	"github.com/datatogether/core"
)

type CustomCrawls int

type CustomCrawlsGetParams struct {
	Id string
}

func (u *CustomCrawls) Get(p *CustomCrawlsGetParams, res *core.CustomCrawl) (err error) {
	url := &core.CustomCrawl{
		Id: p.Id,
	}
	err = url.Read(store)
	if err != nil {
		return err
	}

	*res = *url
	return nil
}

type CustomCrawlsListParams struct {
	OrderBy string
	Limit   int
	Offset  int
}

func (u *CustomCrawls) List(p *CustomCrawlsListParams, res *[]*core.CustomCrawl) (err error) {
	urls, err := core.ListCustomCrawls(store, p.Limit, p.Offset)
	if err != nil {
		return err
	}
	*res = urls
	return nil
}

func (u *CustomCrawls) Save(model *core.CustomCrawl, res *core.CustomCrawl) (err error) {
	err = model.Save(store)
	if err != nil {
		return err
	}

	*res = *model
	return nil
}

func (u *CustomCrawls) Delete(model *core.CustomCrawl, res *core.CustomCrawl) (err error) {
	err = model.Delete(store)
	if err != nil {
		return err
	}

	*res = *model
	return nil
}
