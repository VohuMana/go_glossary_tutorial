package main

import (
	"fmt"
	"log"
	"flag"
)

func main() {
	fmt.Println("Hello LiveEdu! :D")

	var (
		wordToDefine = flag.String("word", "ace", "Enter the word you would like to define.")
	)
	flag.Parse()

	var dictionaryApi Dictionary
	dictionaryApi = CreatePearsonDictionary()

	definitions, err := dictionaryApi.DefineWord(*wordToDefine)
	if err != nil {
		log.Fatal(err)
	}
	
	if len(definitions) == 0 {
		fmt.Println("No definitions found")
	} else {
		for index, definition := range definitions {
			fmt.Println(fmt.Sprintf("%v - %s", index + 1, definition))
		}
	}
}