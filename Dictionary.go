package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Dictionary interface {
	DefineWord(word string) ([]string, error)
}

type PearsonDictionary struct {
	urlFormat string
}

func CreatePearsonDictionary() *PearsonDictionary {
	return &PearsonDictionary{
		urlFormat: "http://api.pearson.com/v2/dictionaries/ldoce5/entries?headword=%s",
	}
}

func (d *PearsonDictionary) DefineWord(word string) ([]string, error) {
	var err error
	var respBody []byte
	var pearsonResponse PearsonDictionary_API
	definitions := []string{}

	response, err := http.Get(fmt.Sprintf(d.urlFormat, word))

	if err == nil {
		defer response.Body.Close()
		respBody, err = ioutil.ReadAll(response.Body)
	}

	if err == nil {
		err = json.Unmarshal(respBody, &pearsonResponse)
	}

	if err == nil {
		for _,result := range pearsonResponse.Results {
			for _,sense := range result.Senses {
				for _,definition := range sense.Definition {
					definitions = append(definitions, definition)
				}
			}
		}
	}

	return definitions, err
}