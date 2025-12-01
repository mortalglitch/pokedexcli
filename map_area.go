package main

type Area struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
}

type AreaList struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	List     []Area `json:"results"`
}
