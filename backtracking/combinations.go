package main

import "fmt"

func combine(n int, k int) [][]int {
	var results [][]int
	var candidate []int
	var list []int

	for num := 1; num <= n; num++ {
		list = append(list, num)
	}

	backtrack(list, candidate, &results, k, 0)
	return results
}

func backtrack(list []int, candidate []int, results *[][]int, length int, index int) {
	if len(candidate) == length {
		copyCandidate := make([]int, len(candidate))
		copy(copyCandidate, candidate)
		*results = append(*results, copyCandidate)
		return
	}

	for i := 0; i < len(list); i++ {
		num := list[i]
		if len(candidate) > 0 && num <= candidate[len(candidate)-1] {
			continue
		}
		candidate = append(candidate, list[i])               // take
		backtrack(list, candidate, results, length, index+1) // explore
		candidate = candidate[:len(candidate)-1]             // clean up
	}
}

func main() {
	fmt.Println(combine(4, 2)) // [[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
	fmt.Println(combine(1, 1)) // [[1]]
}
