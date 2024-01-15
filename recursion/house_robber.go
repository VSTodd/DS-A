package main

import "fmt"

func rob(nums []int) int {
	memo := make(map[int]int)
	return helper(nums, memo)
}

func helper(nums []int, memo map[int]int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	_, ok := memo[len(nums)]

	if ok {
		return memo[len(nums)]
	}

	optionA := helper(nums[1:], memo)
	optionB := nums[0] + helper(nums[2:], memo)

	if optionA >= optionB {
		memo[len(nums)] = optionA
		return optionA
	} else {
		memo[len(nums)] = optionB
		return optionB
	}
}

func main() {
	test1 := []int{1, 2, 3, 1}
	test2 := []int{2, 7, 9, 3, 1}
	test3 := []int{1, 2, 1, 1}

	fmt.Println(rob(test1)) // 4
	fmt.Println(rob(test2)) // 12
	fmt.Println(rob(test3)) // 3
}
