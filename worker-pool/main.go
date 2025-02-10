package main

import "fmt"

func main() {
	inputs := [][]int{
		{2, 3, 4, 5},
	}

	for _, input := range inputs {
		output := work(input)
		fmt.Printf("input:%+v\noutput:%+v\n", input, output)
	}
}

func work(inputs []int) map[int]int {
	results := make(map[int]int)

	workerNum := 3
	jobNum := len(inputs)

	jobChan := make(chan Job, jobNum)
	resultsChan := make(chan Result, jobNum)

	for i := 0; i < workerNum; i++ {
		go worker(jobChan, resultsChan)
	}

	for j := 0; j < jobNum; j++ {
		jobChan <- Job{inputs[j]}
	}
	close(jobChan)

	for k := 0; k < jobNum; k++ {
		result := <-resultsChan
		results[result.Input] = result.Output
	}
	return results
}

type Job struct {
	Input int
}

type Result struct {
	Input  int
	Output int
}

func worker(jobs <-chan Job, results chan<- Result) {
	for job := range jobs {

		square := job.Input * job.Input
		results <- Result{Input: job.Input, Output: square}
	}
}
