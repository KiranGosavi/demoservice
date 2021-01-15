package webCrawler

import (
	"demoservice/entity"
	"demoservice/webCrawler/colly"
)

type Crawler interface {
	GetWebSiteInfo(website *entity.Website) (*entity.WebsiteDetails, error)
}

var Get = func() Crawler {
	return &colly.Impl{}
}

