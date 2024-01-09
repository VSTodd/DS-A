/*
## Problem
- Given a matrix (presented in an array of array), return a single level array of the order going in an spiral
- m is the length of the array (how many rows in the matrix)
- n is the length of the subarrays (how many columns in the matrix)

## Examples
[[1,2,3],[4,5,6],[7,8,9]] returns [1,2,3,6,9,8,7,4,5]
[[1,2,3,4],[5,6,7,8],[9,10,11,12]] returns [1,2,3,4,8,12,11,10,9,5,6,7]

## Data
- Input: nested array
- Output: array

## Algorithm
- Initialize empty array of integers "order"

- create a loop, that breaks when "matrix" is empty
    - remove the first subarray from the array
        - add every value from subarray to order
    - iterate over each value of matrix except for the last, popping the last value from each subarray and adding to order
    - remove the last subarray from the array
        - add every value from subarray in reverse order
    - iterate over each value of matrix in reverse order, popping the first value from each subarray and adding to order

    - iterate over each subarray; removing any empty subarray
    - if matrix is empty, break
*/
package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var order []int
	var holding []int
	for {

		// Top
		holding, matrix = matrix[0], matrix[1:]
		for _, num := range holding {
			order = append(order, num)
		}

		if len(matrix) == 0 {
			break
		}

		// Down
		sublasti := len(matrix[0]) - 1

		if sublasti < 0 {
			break
		}

		for i := 0; i < len(matrix)-1; i++ {
			order = append(order, matrix[i][sublasti])
			matrix[i] = matrix[i][0:sublasti]
		}

		// Right
		lasti := len(matrix) - 1

		holding, matrix = matrix[lasti], matrix[0:lasti]

		for i := len(holding) - 1; i >= 0; i-- {
			order = append(order, holding[i])
		}

		if len(matrix) == 0 {
			break
		}

		if len(matrix[0]) == 0 {
			break
		}

		// Up

		for i := len(matrix) - 1; i >= 0; i-- {
			order = append(order, matrix[i][0])
			matrix[i] = matrix[i][1:]
		}

	}
	return order
}

func main() {
	test1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	test2 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(test1)) // [1,2,3,6,9,8,7,4,5]
	fmt.Println(spiralOrder(test2)) // [1,2,3,4,8,12,11,10,9,5,6,7]
}
