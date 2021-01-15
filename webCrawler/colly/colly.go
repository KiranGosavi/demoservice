package colly

import (
	"bytes"
	"github.com/KiranGosavi/demoservice/entity"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/url"
	"strings"
)

type Impl struct{}

func (i *Impl) GetWebSiteInfo(website *entity.Website) (*entity.WebsiteDetails, error) {
	u, retErr := url.Parse(website.URL)
	if retErr != nil {
		return nil, retErr
	}
	domain := u.Hostname()

	var retInfo entity.WebsiteDetails

	retInfo.URL = website.URL

	c := colly.NewCollector(
		colly.MaxDepth(1), // only the current page. no children pages.
	)

	// Callback for when a scraped page contains an article element
	c.OnHTML("title", func(e *colly.HTMLElement) {
		retInfo.Title = e.Text
	})

	c.OnError(func(response *colly.Response, err error) {
		log.Println("could not get the website details")
		retErr = err
	})

	// Callback for links on scraped pages
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Extract the linked URL from the anchor tag
		link := e.Attr("href")
		if strings.Contains(link, domain) || strings.HasPrefix(link, "/") {
			retInfo.InternalLinks++
		} else {
			retInfo.ExternalLinks++
		}
	})

	// check the login form
	// Callback for links on scraped pages
	c.OnHTML("input", func(e *colly.HTMLElement) {
		if e.Attr("type") == "password" &&
			strings.Contains(e.Attr("class"), "form-control") {
			retInfo.LoginPagePresent = true
		}
	})

	// c.OnHtml only returns text tokens so to get doctype token use other NewTokenizer
	// this is to get the html version !DOCTYPE
	c.OnResponse(func(response *colly.Response) {
		tokenizer := html.NewTokenizer(bytes.NewReader(response.Body))
		for {
			tt := tokenizer.Next()
			t := tokenizer.Token()
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			switch tt {
			case html.ErrorToken:
				retErr = errors.New("could not tokenize response")
				break
			case html.DoctypeToken:
				data := strings.TrimSpace(t.Data)
				retInfo.HTMLVersion = getVersion(data)
				break
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	retErr = c.Visit(website.URL)
	if retErr != nil {
		return nil, retErr
	}
	return &retInfo, nil
}

//A function to get version of a html page.
func getVersion(data string) string {
	if data == "html" {
		return "Html 5"
	}
	if strings.Contains(data, "HTML 4.01") {
		return "HTML 4.01 Strict/Transactional/Frameset"
	}
	if strings.Contains(data, "XHTML 1.0") {
		return "XHTML 1.0 Strict/Transactional/Frameset"
	}
	return ""
}
