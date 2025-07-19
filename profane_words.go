package main

import "strings"

func validateProfaneWords(body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ")

	for i, w := range words {
		loweredWord := strings.ToLower(w)
		if _, ok := badWords[loweredWord]; ok {
			words[i] = "****"
		}
	}

	cleanedText := strings.Join(words, " ")
	return cleanedText
}
