package entity

//A structure to hold url received in a request
type Website struct {
	URL string `json:"website_url"`
}

//A structure to hold the output of webcrawler data
type WebsiteDetails struct {
	URL               string `json:"website_url"`
	Title             string `json:"title"`
	HTMLVersion       string `json:"html_version"`
	InternalLinks     int    `json:"internal_links"`
	ExternalLinks     int    `json:"external_links"`
	InaccessibleLinks int    `json:"inaccessible_links"`
	LoginPagePresent  bool   `json:"login_page"`
}
