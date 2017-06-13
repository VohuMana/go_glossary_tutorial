package main

import (
	"fmt"
	"log"
	"flag"
	"net/http"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello LiveEdu! :D")

	var (
		wordToDefine = flag.String("word", "ace", "Enter the word you would like to define.")
	)
	flag.Parse()

	response, err := http.Get(fmt.Sprintf("http://api.pearson.com/v2/dictionaries/ldoce5/entries?headword=%s", *wordToDefine))
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))
}