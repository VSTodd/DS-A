/*
## Problem
- Given a roman numeral, convert to an integer
- The length of the roman numeral will be between 1 and 15 characters
- The value of the roman numberal will be between 1 and 3999

Considerations
- While in general roman numerals are written largest to smallest, there are six instances where a smaller number occurs before a larger, indicating subtraction
    - I (1) before V(5) or X(10)
    - X (10) before L(50) or C(100)
    - C (100) before D(500) or M(1000)

### Examples
- "III" returns 3
- "LVIII" return 58
- "MCMXCIV" returns 1000 + 900 + 90 + 4 = 1994

### Data Structures
- Input: string
- Output: integer

- Have a "key" which is a map structure
- Big O time complexity - O(N)

### Algorithm
- Iterate over each rune in the string, adding it to the total value
- If the roman value of the rune is greater than the one that proceeds it, subtract the value of the preceeding rune x2 from the total value
*/

package main

import "fmt"

func romanToInt(s string) int {
	roman := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var total int

	for i, v := range s {
		curval := roman[string(v)]
		preval := 1001 // assigning to a value higher than anything in the key, so automatically no subtraction occurs if first character

		if i >= 1 {
			preval = roman[string(s[i-1])]
		}

		total += curval
		if preval < curval {
			total -= (2 * preval)
		}
	}

	return total
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
