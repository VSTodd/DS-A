package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	memo := make(map[[2]int]int)
	return helper(triangle, 0, 0, memo)
}

func helper(triangle [][]int, row int, i int, memo map[[2]int]int) int {
	if row >= len(triangle) || i >= len(triangle[row]) {
		return 0
	}

	val, ok := memo[[2]int{row, i}]

	if ok {
		return val
	}

	memo[[2]int{row, i}] = triangle[row][i] + min(helper(triangle, row+1, i, memo), helper(triangle, row+1, i+1, memo))

	return memo[[2]int{row, i}]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	sub1a := []int{2}
	sub1b := []int{3, 4}
	sub1c := []int{6, 5, 7}
	sub1d := []int{4, 1, 8, 3}
	test1 := [][]int{sub1a, sub1b, sub1c, sub1d} // 11

	sub2a := []int{-1}
	sub2b := []int{2, 3}
	sub2c := []int{1, -1, -3}
	test2 := [][]int{sub2a, sub2b, sub2c} // -1

	fmt.Println(minimumTotal(test1))
	fmt.Println(minimumTotal(test2))
}
