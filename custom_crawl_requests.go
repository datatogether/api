package main

import (
	"github.com/datatogether/archive"
)

type CustomCrawls int

type CustomCrawlsGetParams struct {
	Id string
}

func (u *CustomCrawls) Get(p *CustomCrawlsGetParams, res *archive.CustomCrawl) (err error) {
	url := &archive.CustomCrawl{
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

func (u *CustomCrawls) List(p *CustomCrawlsListParams, res *[]*archive.CustomCrawl) (err error) {
	urls, err := archive.ListCustomCrawls(store, p.Limit, p.Offset)
	if err != nil {
		return err
	}
	*res = urls
	return nil
}

func (u *CustomCrawls) Save(model *archive.CustomCrawl, res *archive.CustomCrawl) (err error) {
	err = model.Save(store)
	if err != nil {
		return err
	}

	*res = *model
	return nil
}

func (u *CustomCrawls) Delete(model *archive.CustomCrawl, res *archive.CustomCrawl) (err error) {
	err = model.Delete(store)
	if err != nil {
		return err
	}

	*res = *model
	return nil
}
