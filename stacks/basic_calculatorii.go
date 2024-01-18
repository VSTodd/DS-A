package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	regex := regexp.MustCompile("[+-/*]")
	nums := regex.Split(s, -1)
	operands := regex.FindAllString(s, -1)

	first, _ := strconv.Atoi(nums[0])
	stack := []int{first}

	for i := 0; i < len(operands); i++ {
		current, _ := strconv.Atoi(nums[i+1])
		prev := stack[len(stack)-1]

		switch operands[i] {
		case "/":
			stack = stack[:len(stack)-1]
			stack = append(stack, prev/current)
		case "*":
			stack = stack[:len(stack)-1]
			stack = append(stack, current*prev)
		case "-":
			stack = append(stack, 0-current)
		case "+":
			stack = append(stack, current)
		}
	}

	total := 0

	for i := len(stack) - 1; i >= 0; i-- {
		total = total + stack[i]
	}

	return total
}

func main() {
	fmt.Println(calculate("3+2*2"))     // 7
	fmt.Println(calculate(" 3/2 "))     // 1
	fmt.Println(calculate(" 3+5 / 2 ")) // 5
}
