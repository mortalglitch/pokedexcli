package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var CurrentMapAreaList AreaList

func startMapData() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &CurrentMapAreaList); err != nil{
		return err
	}

	return nil
}

func nextMapData() error {
	newURL := CurrentMapAreaList.Next
	if newURL != "null" {
		res, err := http.Get(newURL)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(body, &CurrentMapAreaList); err != nil{
			return err
		}
	} else {
		fmt.Println("No additional map data")
	}
	return nil
}

func previousMapData() error {
	newURL := CurrentMapAreaList.Previous
	if newURL != "null" {
		res, err := http.Get(newURL)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(body, &CurrentMapAreaList); err != nil{
			return err
		}
	} else {
		fmt.Println("No previous map data")
	}
	return nil
}


