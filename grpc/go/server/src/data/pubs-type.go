package data

import (
	"encoding/json"
	"io/ioutil"
)

//Pub datastructure
type Pub struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Cite     string `json:"cite"`
	Link     string `json:"link"`
	Slides   string `json:"slides"`
	Abstract string `json:"abstract"`
}

//LoadPubs from Resource File
func LoadPubs() ([]Pub, error) {
	pubs, err := ioutil.ReadFile("./resources/pubs.json")
	if err != nil {
		return nil, err
	}

	var pubDB []Pub

	if err := json.Unmarshal(pubs, &pubDB); err != nil {
		return nil, err
	}

	return pubDB, nil
}
