package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(str string) bool {
	regStr := strings.Join(regexp.MustCompile("([[:alpha:]]|\\d)*").FindAllString(strings.ToLower(str), -1), "")
	return processString(regStr, 0, len(regStr)-1)
}

func processString(str string, start int, end int) bool {
	if end-start < 1 {
		return true
	}
	return str[start] == str[end] && processString(str, start+1, end-1)
}

func main() {
	input1 := "A man, a plan, a canal: Panama"
	input2 := "race a car"
	input3 := " "

	fmt.Println(isPalindrome(input1)) // true
	fmt.Println(isPalindrome(input2)) // false
	fmt.Println(isPalindrome(input3)) // true
}
