/*
### Problem
Find if there are two numbers in the array in which one is equal to double the other. Return true if so, otherwise false.

## Dummy solution
- Nested loops O(N2)
- Goal: O(N logN)

## Data
- With each index, loop against a slice of the array past the index O(N)
*/

package main

import "fmt"

func checkIfExist(arr []int) bool {
	exist := false
out:
	for i := 0; i < len(arr)-1; i++ {
		slice := arr[(i + 1):]
		for _, num := range slice {
			if 2*num == arr[i] || 2*arr[i] == num {
				exist = true
				break out
			}
		}
	}

	return exist
}

func main() {
	test1 := []int{10, 2, 5, 3}
	test2 := []int{3, 1, 7, 11}
	test3 := []int{7, 1, 14, 11}
	fmt.Println(checkIfExist(test1)) // true
	fmt.Println(checkIfExist(test2)) // false
	fmt.Println(checkIfExist(test3)) // true
}
