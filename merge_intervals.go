/*
## Problem:
- Given an array of intervals merge all overlapping intervals, returning an array of the non-overlapping intervals that cover all the intervals in the input
- Intervals all have a length of two, and overlap if the second val of an interval is equal two or greater than the first val of the next interval

## Examples
- [[1,3],[2,6],[8,10],[15,18]] returns [[1,6],[8,10],[15,18]]
- [[1,4],[4,5]] returns [[1,5]]

## Data
- Input: [][]int
- Output: [][]int

- Modify the original array

## Algorithm:
- Utulize a for loop within a regular loop:
    - Initialize empty array, merged
    - Initialize variable nochanges, which is true
        - If the value at the second index is equal to or greater than first index of next value
            - set nochanges to false
            - add subarray of first index of first value and second index of second value to merged
            - add 1 to index
        - Else
            - add current value as it
    - Reassign intervals to merged
    - If nochanges is true, break out of the loop
- Return intervals

*/

package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	nochanges := true
	var merged [][]int
	var empty [][]int
	for {
		nochanges = true
		merged = empty
		for i := 0; i < len(intervals); i++ {

			if i < len(intervals)-1 && intervals[i][1] >= intervals[i+1][0] {
				nochanges = false
				first := intervals[i][0]
				if intervals[i+1][0] < first {
					first = intervals[i+1][0]
				}
				second := intervals[i+1][1]
				if intervals[i][1] > second {
					second = intervals[i][1]
				}
				newsub := []int{first, second}
				merged = append(merged, newsub)
				i++
			} else {
				merged = append(merged, intervals[i])
			}
		}

		intervals = merged

		if nochanges {
			break
		}
	}

	return intervals
}

func main() {
	test1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	test2 := [][]int{{1, 4}, {4, 5}}

	fmt.Println(merge(test1)) // [[1,6],[8,10],[15,18]]
	fmt.Println(merge(test2)) // [[1,5]])
}
