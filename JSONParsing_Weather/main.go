package main

import (
	"fmt"
	"log"
	"flag"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func main() {
	fmt.Println("Hello LiveEdu! Session 1.5 :D")

	var (
		lat = flag.Float64("lat", 51.51, "The latitude of the city you would like the weather for")
		lon = flag.Float64("lon", -0.13, "The longitude of the city you would like the weather for")
		unitType = flag.Int("units", 0, "The units you would like to see 0 is for C (metric), 1 is for F (imperial), and any other number is for K (standard)")
	)
	flag.Parse()

	var unitsName string
	var tempAbbrv string
	switch *unitType {
		case 0:
			unitsName = "metric"
			tempAbbrv = "C"

		case 1:
			unitsName = "imperial"
			tempAbbrv = "F"

		default:
			unitsName = ""
			tempAbbrv = "K"
	}

	// API key: c9fd32706067b93ee688e4e06758f700
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=c9fd32706067b93ee688e4e06758f700&units=%s", *lat, *lon, unitsName)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weatherAPI OpenWeatherApi_RootObject
	err = json.Unmarshal(respBody, &weatherAPI)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The weather for %v, %v is %v %s.  You can expect the low to be %v %s and the high to be %v %s",
		 weatherAPI.Name, 
		 weatherAPI.Sys.Country, 
		 weatherAPI.Main.Temp, 
		 tempAbbrv,
		 weatherAPI.Main.Temp_min, 
		 tempAbbrv,
		 weatherAPI.Main.Temp_max,
		 tempAbbrv)
}