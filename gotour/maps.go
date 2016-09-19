package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	splitted_words := strings.Fields(s)
	wc_map := make(map[string]int)

	for word := 0; word < len(splitted_words); word++ {
		wc_map[splitted_words[word]]++
	}

	return wc_map
}

func main() {
	wc.Test(WordCount)
}
