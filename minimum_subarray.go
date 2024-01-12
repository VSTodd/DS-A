package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	a, r := 0, 0
	shortest := len(nums) + 1
	currentTotal := nums[0]
	for {
		if r == (len(nums)-1) && currentTotal < target {
			break
		}

		if currentTotal < target {
			r++
			currentTotal += nums[r]
		} else {
			length := r - a + 1
			if length < shortest {
				shortest = length
			}
			currentTotal = currentTotal - nums[a]
			a++
		}
	}

	if shortest > len(nums) {
		return 0
	} else {
		return shortest
	}
}

func main() {
	test1 := []int{2, 3, 1, 2, 4, 3}
	test2 := []int{1, 4, 4}
	test3 := []int{1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(minSubArrayLen(7, test1))  // 2
	fmt.Println(minSubArrayLen(4, test2))  // 1
	fmt.Println(minSubArrayLen(11, test3)) // 0
}
