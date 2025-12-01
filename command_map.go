package main

import (
	"fmt"
)

func getCurrentMap() error {
	if CurrentMapAreaList.Next == ""{
		startMapData()
	} else {
		nextMapData()
	}

	for _, i := range CurrentMapAreaList.List {
		fmt.Println(i.Name)
	}

	return nil
}

func getPreviousMap() error {
	if CurrentMapAreaList.Previous == ""{
		return fmt.Errorf("No previous map data")
	} else {
		previousMapData()
	}

	for _, i := range CurrentMapAreaList.List {
		fmt.Println(i.Name)
	}

	return nil
}

