package main

import (
	"fmt"
	"sort"
)

func main() {
	inputs := [][]Person{
		{{"Alice", 30}, {"Charlie", 35}, {"Bob", 25}},
	}

	for _, input := range inputs {
		output := make([]Person, len(input))
		_ = copy(output, input)

		SortPersonsByAge(output)
		fmt.Printf("Sort By Age\ninput: %+v\noutput: %+v\n", input, output)

		output = make([]Person, len(input))
		_ = copy(output, input)

		SortPersonsByName(output)
		fmt.Printf("Sort By Name\ninput: %+v\noutput: %+v\n", input, output)
	}
}

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func SortPersonsByAge(persons []Person) {
	sort.Sort(ByAge(persons))
}

type ByName []Person

func (b ByName) Len() int {
	return len(b)
}

func (b ByName) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}

func (b ByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func SortPersonsByName(persons []Person) {
	sort.Sort(ByName(persons))
}
