package main

import (
	"encoding/json"
	"io/ioutil"
)

//PubDbRecord datastructure
type PubDbRecord struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Cite     string `json:"cite"`
	Link     string `json:"link"`
	Slides   string `json:"slides"`
	Abstract string `json:"abstract"`
}

//LoadPubs from Resource File
func LoadPubs() ([]PubDbRecord, error) {
	pubs, err := ioutil.ReadFile("./pubs.json")
	if err != nil {
		return nil, err
	}

	var pubDB []PubDbRecord

	if err := json.Unmarshal(pubs, &pubDB); err != nil {
		return nil, err
	}

	return pubDB, nil
}
