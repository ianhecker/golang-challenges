# Implement a Worker Pool

Time Limit: 60 minutes

## Instructions

Design a worker pool in Golang that processes a set of tasks concurrently.

Each task can be a simple operation, such as calculating the square of a number.

Create a fixed number of worker goroutines to handle tasks concurrently.
Use a channel to send tasks to workers and another channel to collect results.
Ensure proper shutdown of the worker pool after all tasks are processed.

Example Task and Output:
---
Task Input: []int{2, 3, 4, 5}
Task Operation: Square each number
Output: map[int]int{2: 4, 3: 9, 4: 16, 5: 25}
---

Reuirements:

The number of workers should be configurable.
Tasks should be distributed evenly among workers.
Implement a way to signal workers to stop once all tasks have been processed (e.g., close channels or use a sync.WaitGroup).

Notes:

Use goroutines and channels to achieve concurrent task execution.
Handle errors or panics gracefully within the workers.
Consider using buffered channels to optimize performance in high-load scenarios.
