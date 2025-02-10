package main

import "fmt"

func main() {

	inputs := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{10, -5, 15, 0},
		{},
	}

	for _, input := range inputs {

		done := make(chan int, 4)

		i1 := input[len(input)/2:]
		i2 := input[:len(input)/2]

		go sumSlice(done, i1...)
		go sumSlice(done, i2...)

		output := 0
		for i := 0; i < 2; i++ {
			select {
			case sum := <-done:
				output += sum
			}
		}
		fmt.Printf("input: %+v\noutput: %d\n", input, output)
	}
}

func sumSlice(done chan int, s ...int) {
	sum := 0
	for _, n := range s {
		sum += n
	}
	done <- sum
}
