package server

import (
	"io/ioutil"
	"encoding/json"
)



func PubDB() ([]Pub, error) {
	pubs, err := ioutil.ReadFile("./resources/pubs.json")
	if err != nil {
		return nil,err
	}

	var pubDB []Pub

	if err := json.Unmarshal(pubs, &pubDB); err != nil {
		return nil,err
	}

	return pubDB,nil
}
