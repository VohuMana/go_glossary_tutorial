package main

type PearsonDictionary_API struct {
	Status float64 `json:"status"`
	Offset float64 `json:"offset"`
	Limit float64 `json:"limit"`
	Count float64 `json:"count"`
	Total float64 `json:"total"`
	Url string `json:"url"`
	Results []Results `json:"results"`
}

type Results struct {
	Url string `json:"url"`
	Datasets []string `json:"datasets"`
	Headword string `json:"headword"`
	Homnum float64 `json:"homnum"`
	Id string `json:"id"`
	Part_of_speech string `json:"part_of_speech"`
	Senses []Senses `json:"senses"`
}

type Senses struct {
	Definition []string `json:"definition"`
	Examples []Examples `json:"examples"`
	Lexical_unit string `json:"lexical_unit"`
}

type Examples struct {
	Audio []Audio `json:"audio"`
	Text string `json:"text"`
}

type Audio struct {
	Type string `json:"type"`
	Url string `json:"url"`
}

