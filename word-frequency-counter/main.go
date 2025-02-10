package main

import (
	"fmt"
	"strings"
)

func main() {
	inputs := []string{
		"Hello, world! Hello.",
		"Go is great! Go, go, go!",
	}

	for _, input := range inputs {

		output := WordFrequency(input)
		fmt.Printf("%s\n%+v\n", input, output)
	}
}

func WordFrequency(s string) map[string]int {
	tmp := strings.ToLower(s)
	tmp = strings.ReplaceAll(tmp, ",", "")
	tmp = strings.ReplaceAll(tmp, "!", "")
	tmp = strings.ReplaceAll(tmp, ".", "")
	words := strings.Split(tmp, " ")

	frequency := make(map[string]int)

	for _, word := range words {
		count := frequency[word]
		frequency[word] = count + 1
	}

	return frequency
}
