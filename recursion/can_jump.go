package main

import "fmt"

func canJump(nums []int) bool {
	memo := make(map[int]bool)
	memo[0] = true
	return helper(nums, memo)
}

func helper(nums []int, memo map[int]bool) bool {
	for i, num := range nums {
		if memo[i] == false {
			continue
		}

		for j := 1; j <= num && i+j < len(nums); j++ {
			memo[i+j] = true
		}
	}

	return memo[len(nums)-1]
}

func main() {
	test1 := []int{2, 3, 1, 1, 4}
	test2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(test1)) // true
	fmt.Println(canJump(test2)) // false
}
