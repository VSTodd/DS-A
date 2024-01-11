/*
### Problem
- Find the an element where all adjacent elements are smaller than it (it's okay if at front or end of array)

### Examples
Input: nums = [1,2,3,1]
Output: 2

Input: nums = [1,2,1,3,5,6,4]
Output: 5

### Data
- Do a binary search

### Algorithm
- After identifying a mid, if not a peak go shift to any side where the adjacent number is higher
*/

package main

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	} else if len(nums) == 2 {
		if nums[0] >= nums[1] {
			return 0
		} else {
			return 1
		}
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) / 2)
		if mid > 0 && mid < (len(nums)-1) && nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if mid == 0 && nums[mid] > nums[mid+1] {
			return mid
		} else if mid == (len(nums)-1) && nums[mid] > nums[mid-1] {
			return mid
		} else if mid > 0 && nums[mid] < nums[mid-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	test1 := []int{1, 2, 3, 1}          // 2
	test2 := []int{1, 2, 1, 3, 5, 6, 4} // 5
	fmt.Println(findPeakElement(test1))
	fmt.Println(findPeakElement(test2))
}
