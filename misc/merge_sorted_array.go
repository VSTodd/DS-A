/*
## Problem
- Two integer arrays (nums1 & nums2), in ascending order
- m (nums1)& n(nums2), represent how many elements are in each array
- Merge the two arrays into a single array in ascending order
- The final array should be stored inside nums1
- The length of nums1 is m + n
- At least one array will have a valid value
- 0 could be a valid value

## Examples:
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]

Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]

Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]

## Data
- Slow approach O(n logn):
  - Replacing all 0s in nums1 with numbers in nums2, then sort

- Better approach O(n + m):

- Have three pointers, "indexTotal" at the end of nums1, index1 at nums1(m - 1) and index2 at nums2(n-1)
  - compare index1 and index2, adding the larger one to indexTotal
  - subtract indexTotal by one and the used index by 1
  - break when indexTotal is less than 0
*/
package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	indexTotal := len(nums1) - 1
	index1 := m - 1
	index2 := n - 1

	for indexTotal >= 0 {
		if index2 < 0 || index1 >= 0 && nums1[index1] >= nums2[index2] {
			nums1[indexTotal] = nums1[index1]
			index1--
		} else if index1 < 0 || nums1[index1] < nums2[index2] {
			nums1[indexTotal] = nums2[index2]
			index2--
		}
		indexTotal--
	}

	return nums1
}

func main() {
	test1a := []int{1, 2, 3, 0, 0, 0}
	test2a := []int{2, 5, 6}
	ma := 3
	na := 3
	fmt.Println(merge(test1a, ma, test2a, na)) // [1,2,2,3,5,6]

	test1b := []int{1}
	test2b := []int{0}
	mb := 1
	nb := 0
	fmt.Println(merge(test1b, mb, test2b, nb)) // [1]

	test1c := []int{0}
	test2c := []int{1}
	mc := 0
	nc := 1
	fmt.Println(merge(test1c, mc, test2c, nc)) // [1]
}
