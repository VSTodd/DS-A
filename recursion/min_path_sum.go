package main

import "fmt"

func minPathSum(grid [][]int) int {
	memo := make(map[[2]int]int)
	return helper(grid, 0, 0, memo)
}

func helper(grid [][]int, row int, column int, memo map[[2]int]int) int {
	_, ok := memo[[2]int{row, column}]
	if ok {
		return memo[[2]int{row, column}]
	}

	if row >= len(grid)-1 && column >= len(grid[0])-1 {
		memo[[2]int{row, column}] = grid[row][column]
		return memo[[2]int{row, column}]
	} else if row >= len(grid)-1 {
		memo[[2]int{row, column}] = grid[row][column] + helper(grid, row, column+1, memo)
		return memo[[2]int{row, column}]
	} else if column >= len(grid[0])-1 {
		memo[[2]int{row, column}] = grid[row][column] + helper(grid, row+1, column, memo)
		return memo[[2]int{row, column}]
	}

	option1 := helper(grid, row, column+1, memo)
	option2 := helper(grid, row+1, column, memo)

	memo[[2]int{row, column + 1}] = option1
	memo[[2]int{row + 1, column}] = option2

	return grid[row][column] + min(option1, option2)
}

func main() {
	sub1a := []int{1, 3, 1}
	sub1b := []int{1, 5, 1}
	sub1c := []int{4, 2, 1}
	arr1 := [][]int{sub1a, sub1b, sub1c}

	fmt.Println(minPathSum(arr1)) // 7

	sub2a := []int{1, 2, 3}
	sub2b := []int{4, 5, 6}
	arr2 := [][]int{sub2a, sub2b}

	fmt.Println(minPathSum(arr2)) // 12
}
