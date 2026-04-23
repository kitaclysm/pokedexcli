package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var words []string
	fields := strings.Fields(text)
	if len(fields) == 0 {
		return []string{}
	}
	for _, field := range fields {
		field = strings.ToLower(field)
		words = append(words, field)
	}
	return words
}