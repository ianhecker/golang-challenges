# Puzzle 1: Find the Maximum Element

Time Limit: 20 minutes

## Instructions

Write a generic function Max that accepts a slice of any ordered type (e.g.,
integers, floats, strings) and returns the largest element in the slice.

Example Output:
---
Max([]int{1, 3, 2}) -> 3
Max([]float64{1.5, 3.4, 2.2}) -> 3.4
Max([]string{"apple", "orange", "banana"}) -> "orange"
---

# Requirements:

1) The function should use generics with a type constraint for ordered types.
2) Handle slices of different types (int, float64, string).

Optional: Handle the case where the slice is empty.
