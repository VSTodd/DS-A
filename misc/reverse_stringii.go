/*
## Problem:
- When given a string (s) and an integer(k), reverse the first k charactes, every 2k characters
  - If there are fewer than k characters left, reverse all of them
  - If there are less than 2k characters left, but more than/equal to k characters left,
    reverse the first k characters and leave the others be
  - Problem will consist of strings longer than 1, made up only of lowercase letters

## Examples:
Input: s = "abcdefg", k = 2
Output: "bacdfeg"

Input: s = "abcd", k = 2
Output: "bacd"

## Data/Algorithm:
- initialize string, reversed
- initialize boolean swap to true
- initialize index to 0

- utilize a for loop
  - if swap is true
    - take substring
    - reverse it
    add to reverse
  - else
    add substring to reversed
  - change swaps value
  - add i by k
  - break if k is greater than string

- return string
*/

package main

import "fmt"

func reverseStr(s string, k int) string {
	if len(s) <= 1 {
		return s
	}

	var reversed string
	swap := true
	lastIndex := len(s) - 1
	subStringStart := 0
	subStringEnd := findEnd(subStringStart+k, lastIndex)

	for subStringStart <= lastIndex {
		if swap {
			substring := findReverse(s[subStringStart:subStringEnd])
			reversed = reversed + substring
		} else {
			reversed = reversed + s[subStringStart:subStringEnd]
		}

		swap = !swap
		subStringStart = subStringStart + k
		subStringEnd = findEnd(subStringStart+k, lastIndex)
	}

	return reversed
}

func findReverse(sub string) string {
	var revSub string
	for i := len(sub) - 1; i >= 0; i-- {
		revSub = revSub + string(sub[i])
	}

	return revSub
}

func findEnd(end int, last int) int {
	if end > last {
		return last + 1
	} else {
		return end
	}
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2)) // bacdfeg
	fmt.Println(reverseStr("abcd", 2))    // bacd
}
