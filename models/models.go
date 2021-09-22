package models

type Universities struct{
	Name string `json:"name"`
	Abbrv string `json:"abbreviation"`
	WebsiteLink string `json:"websitelink"`
}

type NewUniversities struct{
	Name string `json:"name"`
	Abbrv string `json:"abbreviation"`
	WebsiteLink string `json:"websitelink"`
}