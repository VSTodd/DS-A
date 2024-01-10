/*
## Problem
- Return all triple values that do not have equal indexes that equal to 0
- All arrays are three or larger
- Subarrays must be in ascending order
- Input: array
- Output: array of arrays

## Examples
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Input: nums = [0,1,1]
Output: []

Input: nums = [0,0,0]
Output: [[0,0,0]]

## Data
- Create a hash for all values, so you can calculate what the anchor and pointer need and check the hash for it
  - Do need a protective clause so that you're grabbing from different indexes

## Algorithm
- Create hash, with the value of each number being the number of times it occurs in the array
- initiate array of arrays triplets
- Initialize anchor to 0, and runner to 1
- For loop
  - add value and anchor together and calulcate what number they'd need to reach zero
  - initialize variable count to one
  - add one to count if the number is the same as the value at anchor or runner (two if both)
  - check if has contains the value, equal to or more than count, if so:
      - extract to sort function, that returns sorted array
      - if triplets does not contain the same array, add it to triplets
  - add 1 to runner
  - if runner is equal to the length of nums
    - add anchor by one
    - runner is equal to anchor plus 1
  - if anchor is equal to length of nums minus two
    - break
- return triplets
*/

package main

import (
	"fmt"
	"reflect"
	"slices"
)

func threeSum(nums []int) [][]int {
	mapped := make(map[int]int)
	for _, num := range nums {
		if mapped[num] >= 1 {
			mapped[num] += 1
		} else {
			mapped[num] = 1
		}
	}

	var triplets [][]int
	a, r := 0, 1

	for {
		value := -(nums[a] + nums[r])
		count := 1

		if value == nums[a] {
			count += 1
		}

		if value == nums[r] {
			count += 1
		}

		if mapped[value] >= count {
			subarray := order(nums[a], nums[r], value)
			isunique := unique(triplets, subarray)

			if isunique {
				triplets = append(triplets, subarray)
			}

		}

		r++

		if r == len(nums) {
			a++
			r = a + 1
		}

		if a == len(nums)-2 {
			break
		}
	}
	return triplets
}

func order(num1 int, num2 int, num3 int) []int {
	arr := []int{num1, num2, num3}
	slices.Sort(arr)
	return arr
}

func unique(triplets [][]int, sub []int) bool {
	uniqueval := true
	for _, arr := range triplets {
		if reflect.DeepEqual(arr, sub) {
			uniqueval = false
			break
		}
	}

	return uniqueval
}

func main() {
	test1 := []int{-1, 0, 1, 2, -1, -4} // [[-1,-1,2],[-1,0,1]]
	test2 := []int{0, 1, 1}             // []
	test3 := []int{0, 0, 0}             // [[0,0,0]]
	fmt.Println(threeSum(test1))
	fmt.Println(threeSum(test2))
	fmt.Println(threeSum(test3))
}
