package sprint
/*
Convert Any to Decimal

Learning Objectives:
*	Develop the ability to convert numbers from one base to another using specified characters, enhancing your understanding of numerical representation.
*	Learn how to validate input data, especially the base, to ensure it meets specific criteria or conditions and handle invalid input gracefully.
*	Gain proficiency in manipulating strings to extract and process numeric values represented in different bases.

Instructions
Create a Go function that takes two parameters:
* 	s: a numeric string in a specific base.
* 	base: a string containing unique characters representing the available digits in that base.
	The function should return the integer value of s in the given base. If the base is not valid, it returns 0.
	Here are the validity rules for the base:
* 	The base must contain at least 2 unique characters.
* 	The base should not contain the characters + or -.
	The function only works with valid string numbers that consist of elements present in the base. It does not handle negative numbers.

Allowed packages
For this task you're allowed to use:
*	math
*	strings
*/

import (
	"math"
	"strings"
)

func ConvertAnyToDec(s string, base string) int {
	seen := make(map[rune]bool)
	baseLen := len(base)

	if baseLen < 2 {
		return 0
	}

	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return 0
		}
		seen[r] = true
	}
	for _, r := range s {
		if !strings.ContainsRune(base, r) {
			return 0
		}
	}
	result := 0
	sLen := len(s)
	for i, r := range s {
		value := strings.IndexRune(base, r)
		if value == -1 {
			return 0
		}
		power := sLen - 1 - i
		result += value * int(math.Pow(float64(baseLen), float64(power)))
	}
	return result
}
