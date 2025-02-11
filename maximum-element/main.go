package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	inputInt := []int{1, 3, 2}
	maxInt := Max(inputInt)
	fmt.Printf("input:%+v, output:%v\n", inputInt, maxInt)

	inputFloat := []float64{1.5, 3.4, 2.2}
	maxFloat := Max(inputFloat)
	fmt.Printf("input:%+v, output:%v\n", inputFloat, maxFloat)

	inputString := []string{"apple", "orange", "banana"}
	maxString := Max(inputString)
	fmt.Printf("input:%+v, output:%v\n", inputString, maxString)

}

func Max[T constraints.Ordered](list []T) (max T) {

	for _, item := range list {
		if item > max {
			max = item
		}
	}
	return max
}
