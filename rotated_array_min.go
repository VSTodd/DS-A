/*
## Problem:
- Find the smallest value in a rotated array using binary search
- If the num is smaller than the value on its left, then it's the smallest

## Examples:
Input: nums = [3,4,5,1,2]
Output: 1

Input: nums = [4,5,6,7,0,1,2]
Output: 0

Input: nums = [11,13,15,17]
Output: 11

## Algorithm:
- initialize left and right to the indexes on either end of the array
- if left val is less than right val, return left val

- for loop
  - find the mid index
  - if the element at the mid index is less than the element to its left, return the element
  - if the element at the mid index is greater than the high index element, shift search to the right
  - else, shift search to the left
*/

package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	if len(nums) == 1 {
		return nums[0]
	} else if nums[left] < nums[right] {
		return nums[left]
	}

	for left <= right {
		mid := left + ((right - left) / 2)
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	test1 := []int{3, 4, 5, 1, 2}       // 1
	test2 := []int{4, 5, 6, 7, 0, 1, 2} // 0
	test3 := []int{11, 13, 15, 17}      // 11

	fmt.Println(findMin(test1))
	fmt.Println(findMin(test2))
	fmt.Println(findMin(test3))
}
