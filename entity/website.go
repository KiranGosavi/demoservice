package entity

type Website struct {
	URL string `json:"website_url"`
}
type WebsiteDetails struct {
	URL               string `json:"website_url"`
	Title             string `json:"title"`
	HTMLVersion       string `json:"html_version"`
	InternalLinks     int    `json:"internal_links"`
	ExternalLinks     int    `json:"external_links"`
	InaccessibleLinks int    `json:"inaccessible_links"`
	LoginPagePresent  bool   `json:"login_page"`
}