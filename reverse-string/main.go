package main

import "fmt"

func main() {
	inputs := []string{
		"hello",
		"こんにちは",
		"",
	}

	for _, input := range inputs {
		output := ReverseString(input)
		fmt.Printf("input:'%s'\noutput:'%s'\n", input, output)
	}
}

func ReverseString(s string) string {
	runes := []rune(s)

	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}

	return string(runes)
}
