package main

type Clouds struct {
	All float64 `json:"all"`
}

type OpenWeatherApi_RootObject struct {
	Coord Coord `json:"coord"`
	Weather []Weather `json:"weather"`
	Visibility float64 `json:"visibility"`
	Sys Sys `json:"sys"`
	Name string `json:"name"`
	Base string `json:"base"`
	Main Main `json:"main"`
	Wind Wind `json:"wind"`
	Clouds Clouds `json:"clouds"`
	Dt float64 `json:"dt"`
	Id float64 `json:"id"`
	Cod float64 `json:"cod"`
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Weather struct {
	Description string `json:"description"`
	Icon string `json:"icon"`
	Id float64 `json:"id"`
	Main string `json:"main"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise float64 `json:"sunrise"`
	Sunset float64 `json:"sunset"`
	Type float64 `json:"type"`
	Id float64 `json:"id"`
	Message float64 `json:"message"`
}

type Main struct {
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
	Temp_min float64 `json:"temp_min"`
	Temp_max float64 `json:"temp_max"`
	Temp float64 `json:"temp"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg float64 `json:"deg"`
}

