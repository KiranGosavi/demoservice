package webCrawler

import (
	"github.com/KiranGosavi/demoservice/entity"
	"github.com/KiranGosavi/demoservice/webCrawler/colly"
)

type Crawler interface {
	GetWebSiteInfo(website *entity.Website) (*entity.WebsiteDetails, error)
}

var Get = func() Crawler {
	return &colly.Impl{}
}
