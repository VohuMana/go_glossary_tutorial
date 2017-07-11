package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
	"html"
)

func main() {
	var dictionaryApi Dictionary
	dictionaryApi = CreatePearsonDictionary()

	http.HandleFunc("/api/define", func(w http.ResponseWriter, request *http.Request) {
		var definitions []string
		var splitWords []string
		var jsonObject string
		wordsWithDefinitions := make(map[string][]string)

		words, err := ioutil.ReadAll(request.Body)

		if err == nil {
			splitWords = strings.Split(string(words), " ")
		}

		if err == nil {
			for _, word := range splitWords {
				definitions, err = dictionaryApi.DefineWord(word)
				if err != nil {
					fmt.Println(err)
				}

				wordsWithDefinitions[word] = definitions
			}
		}

		if err == nil {
			jsonBytes, err := json.Marshal(wordsWithDefinitions)

			if err == nil {
				jsonObject = string(jsonBytes)
			}
		}

		if err == nil {
			fmt.Fprintf(w, "%s", jsonObject)
		}
		
		if err != nil {
			fmt.Fprintf(w, "%q", html.EscapeString(err.Error()))
		}
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}