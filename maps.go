package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	word_count := make(map[string]int)
	for _, value := range strings.Fields(s) {
		_, is_mapped := word_count[value]
		if is_mapped {
			word_count[value] += 1
		} else {
			word_count[value] = 1
		}
	}

	return word_count
}

func main() {
	wc.Test(WordCount)
}
