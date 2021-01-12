package handler

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type body struct {
	WebsiteURL string `json:"website_url"`
}
type WebsiteDetails struct {
	URL               string `json:"website_url"`
	Title             string `json:"title"`
	HTMLversion       string `json:"html_version"`
	InternalLinks     int    `json:"internal_links"`
	ExternalLinks     int    `json:"external_links"`
	InaccessibleLinks int    `json:"inaccessible_links"`
	LoginPagePresent  bool   `json:"login_page"`
}

var Websites []WebsiteDetails

func GetWebsiteDetails(w http.ResponseWriter, r *http.Request) {
	var rbody body
	err := json.NewDecoder(r.Body).Decode(&rbody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	wdetails, err := parseURL(rbody.WebsiteURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Websites = append(Websites, wdetails)

	websiteJSON, err := json.Marshal(&wdetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(websiteJSON)

}

func parseURL(url string) (WebsiteDetails, error) {
	//resp, err := http.Get(url)
	//if err != nil {
	//	return WebsiteDetails{},fmt.Errorf("could not parse URL: %v", url)
	//}
	var website WebsiteDetails
	website.URL = url
	details, ok := getCacheDetails(url)
	if ok {
		website = details
	}
	if !ok {
		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(err)
		}

		// use CSS selector found with the browser inspector
		// for each, use index and item
		doc.Find("#main article .entry-title").Each(func(index int, item *goquery.Selection) {
			//title := item.Text()
			website.Title = item.Text()
			linkTag := item.Find("a")
			link, _ := linkTag.Attr("href")
			fmt.Printf("Post #%d: %s - %s\n", index, website.Title, link)
		})
	}
	return website, nil

}

func getCacheDetails(url string) (WebsiteDetails, bool) {
	for _, website := range Websites {
		if website.URL == url {
			return website, true
		}
	}
	return WebsiteDetails{}, false
}
