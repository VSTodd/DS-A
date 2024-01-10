/*
## Problem
Check is a subsequence (a) is present in the given string (t)

## Examples
Input: s = "abc", t = "ahbgdc"
Output: true

Input: s = "axc", t = "ahbgdc"
Output: false


## Data
Input: two string
Output: boolean

- Initialize anchor and runner to zero, moving runner down string t until it's equal to the first letter of string a
- Reassign anchor to runner, and keep moving runner
- Break and return true if anchor is equal to the last character of string a

## Algorithm
- Initialize anchor to -1
- For loop (initialize "runner" to zero, adding by one every time, initialize counter to zero)
  - if runner index at t is equal to counter index at s
    - anchor equals runner
    - counter += 1
  - break out if
    - anchor at t's index is equal to the last letter of s
    - return true
- return false
*/

package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	for runner, anchor := 0, 0; runner < len(t); runner++ {
		if t[runner] == s[anchor] {
			anchor++
		}

		if anchor == len(s) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) // true
	fmt.Println(isSubsequence("axc", "ahbgdc")) // false
}
