package handler

import (
	"encoding/json"
	"github.com/KiranGosavi/demoservice/entity"
	"github.com/KiranGosavi/demoservice/webCrawler"
	"net/http"
)

//handler function to fetch the website details
func GetWebsiteDetails(w http.ResponseWriter, r *http.Request) {
	websiteUrl := r.URL.Query().Get("url")
	if websiteUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	website := entity.Website{
		URL: websiteUrl,
	}

	// get the crawler
	crawler := webCrawler.Get()

	details, err := crawler.GetWebSiteInfo(&website)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	websiteJSON, err := json.Marshal(&details)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(websiteJSON)

}
