package main

import "fmt"

func isValid(s string) bool {
	var stack []string

	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if char == "(" || char == "{" || char == "[" {
			stack = append(stack, char)
			continue
		} else if len(stack) == 0 {
			return false
		}

		open := string(stack[len(stack)-1])

		if (char == "}" && open == "{") || (char == ")" && open == "(") || (char == "]" && open == "[") {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
}
