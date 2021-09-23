package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func JsonToStruct() []Universities {
	jsonFile, _ := os.Open("models.json")
	defer jsonFile.Close()

	jsonList, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	theData := []Universities{}

	_ = json.Unmarshal(jsonList, &theData)

	return theData
}
