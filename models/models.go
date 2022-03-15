package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Universities struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	WebsiteLink  string `json:"website_link"`
}

type NewUniversities struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	WebsiteLink  string `json:"website_link"`
}

func ObjMethod() []Universities {
	jsonList, err := ioutil.ReadFile("models/models.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	var data []Universities

	_ = json.Unmarshal(jsonList, &data)

	return data
}
