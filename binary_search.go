package main

import "fmt"

func binarySearch(array []int, search int) int {
	lowerbound := 0
	upperbound := len(array) - 1
	var midpoint int
	var midpointValue int

	for {
		if lowerbound > upperbound {
			break
		}

		midpoint = (upperbound + lowerbound) / 2
		midpointValue = array[midpoint]

		if midpointValue == search {
			return midpoint
		} else if midpointValue > search {
			upperbound = midpoint - 1
		} else if midpointValue < search {
			lowerbound = midpoint + 1
		}
	}

	return -1
}

func main() {
	test1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	test2 := []int{5, 8, 18, 44, 70, 105, 224, 400}
	fmt.Println(binarySearch(test1, 2))   // 1
	fmt.Println(binarySearch(test1, 10))  // -1
	fmt.Println(binarySearch(test2, 224)) // 6
	fmt.Println(binarySearch(test2, 10))  // -1
}
