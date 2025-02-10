# Concurrent Sum of a Slice

Time Limit: 45 minutes

## Instructions

Implement a program in Golang that computes the sum of all integers in a slice
by dividing the work among multiple goroutines.

Split the slice into segments, either based on the number of available cores or
a fixed number of goroutines.

Create a goroutine for each segment to compute its sum concurrently.
Use channels to collect and combine the partial sums from each goroutine into a
final result.

Example Input and Output:
---
Input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
Output: 55  // Sum of the slice elements

Input: []int{10, -5, 15, 0}
Output: 20
---
Requirements:

The program should handle slices of varying lengths, including edge cases like empty slices.
Aim for efficient concurrent execution by minimizing communication overhead between goroutines.
Consider using the runtime package to determine the number of available CPU cores.

Notes:

Use the sync or runtime packages as needed to optimize performance.
Avoid blocking operations in the main thread while waiting for goroutines to complete.
