/*
## Problem
- Given a matrix of integers, in which each row is in ascending order, and the final integer of each row
	is less than the first integer in the next row:
		- See if a value is present in a matrix
- Solution much be 0(log(N*M)) complexity

## Examples
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false

## Basic Structure
- use binary search as if the matrix is one single, flat array
- once identifying the middle "index", pull out the proper value via an extracted function

## Matrix Algorithm
- m (how many rows) is equal to the length of the array (matrix)
- n (how many columns) is equal to the length of the fist element of matrix
- initialize left to zero
- initialize right to (m * n) - 1

- utilize a typical binary search, identifying the middle value with func matrixValue


## Matrix Value Algorithm
- find the proper row by dividing middle "index" by n, and rounding down
- find the proper index within the row by taking the remainder of middle "index" and n

*/

package main

import (
	"fmt"
	"math"
)

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, (m*n - 1)

	for left <= right {
		mid := left + ((right - left) / 2)
		midValue := matrixValue(mid, n, matrix)
		if midValue == target {
			return true
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func matrixValue(mid int, n int, matrix [][]int) int {
	index1 := int(math.Floor(float64(mid) / float64(n)))
	index2 := mid % n

	return matrix[index1][index2]
}

func main() {
	row1 := []int{1, 3, 5, 7}
	row2 := []int{10, 11, 16, 20}
	row3 := []int{23, 30, 34, 60}
	array := [][]int{row1, row2, row3}
	fmt.Println(searchMatrix(array, 3))  // true
	fmt.Println(searchMatrix(array, 13)) // false
}
