package main

import "strings"

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	cleanText := strings.TrimSpace(lowerText)
	splitString := strings.Fields(cleanText)

	return splitString
}

