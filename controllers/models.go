package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Universities struct {
	Name        string `json:"name"`
	Abbrv       string `json:"abbreviation"`
	WebsiteLink string `json:"websitelink"`
}

// New Universities structure
type NewUniversities struct {
	Name        string `json:"name"`
	Abbrv       string `json:"abbreviation"`
	WebsiteLink string `json:"websitelink"`
	Logo        string `json:"logo"`
}

func ObjMethod() []Universities {
	jsonList, err := ioutil.ReadFile("controllers/models.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	data := []Universities{}

	_ = json.Unmarshal(jsonList, &data)

	return data
}
