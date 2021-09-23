package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Universities struct {
	Name        string `json:"name"`
	Abbrv       string `json:"abbreviation"`
	WebsiteLink string `json:"websitelink"`
}

// New Universities structure
type NewUniversities struct {
	Name        string `json:"name"`
	Abbrv       string `json:"abbrev"`
	WebsiteLink string `json:"weblink"`
}

func ObjMethod() []Universities {
	jsonFile, _ := os.Open("")
	defer jsonFile.Close()

	jsonList, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	data := []Universities{}

	_ = json.Unmarshal(jsonList, &data)

	return data
}
