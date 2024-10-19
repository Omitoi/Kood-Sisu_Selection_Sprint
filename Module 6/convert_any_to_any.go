package sprint

/*
Convert Any to Any

Learning Objectives:
*	Develop the ability to convert numeric values between different bases, understanding the concepts and logic involved in base conversion.
*	Learn to design and implement an efficient algorithm for converting numbers between bases, considering the positional value system.

Instructions
Create a Go function that takes three string arguments: nbr representing a numeric value in a specific base baseFrom,
and baseTo representing the desired base for the returned value.
The function should convert the nbr from baseFrom to baseTo and return the result as a string.

Allowed packages
For this task you're allowed to use:
*	math
*	strings
*/

import (
	"math"
	"strings"
)

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	seen := make(map[rune]bool)
	baseLen := len(baseFrom)
	baseLenTo := len(baseTo)
	var res string

	if baseLen < 2 {
		return ""
	}

	for _, r := range baseFrom {
		if r == '+' || r == '-' || seen[r] {
			return ""
		}
		seen[r] = true
	}
	for _, r := range nbr {
		if !strings.ContainsRune(baseFrom, r) {
			return ""
		}
	}
	result := 0
	sLen := len(nbr)
	for i, r := range nbr {
		value := strings.IndexRune(baseFrom, r)
		if value == -1 {
			return ""
		}
		power := sLen - 1 - i
		result += value * int(math.Pow(float64(baseLen), float64(power)))
	}
	var aa int
	var bb int
	for {
		aa += result % baseLenTo
		result /= baseLenTo
		bb = aa % baseLenTo
		aa /= baseLenTo
		res = string(baseTo[bb]) + res
		if result == 0 {
			break
		}
	}
	return res
}
