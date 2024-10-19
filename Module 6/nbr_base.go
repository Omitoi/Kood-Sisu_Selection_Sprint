package sprint

/*
Number to Base

Learning Objectives:
* 	Develop the ability to convert numbers from one base to another using specified characters, enhancing your understanding of numerical representation.
* 	Learn how to validate input data, especially the base, to ensure it meets specific criteria or conditions and handle invalid input gracefully.
* 	Gain proficiency in manipulating strings to construct representations of numbers in various bases and handle different types of characters.

Instructions
Create a Go function that takes an integer n and a string base as parameters. The function should return the integer n represented in the specified base as a string.
Here are the validity rules for the base:
* 	The base must contain at least 2 unique characters.
* 	The base should not contain the characters + or -.
* 	The function should handle negative numbers as well (see examples in the usage section).
	If the base is not valid, the function should return "NV" (Not Valid).
*/

func NbrBase(n int, base string) string {
	var res string
	seen := make(map[rune]bool)
	baseLen := len(base)
	neg := ""
	if baseLen < 2 {
		return "NV"
	} else if n == 0 {
		return "0"
	}
	if n < 0 {
		n = -n
		neg = "-"
	}
	for i := 0; i < baseLen; i++ {
		if (base[i] >= 'A' && base[i] <= 'Z') || (base[i] >= 'a' && base[i] <= 'z') || (base[i] >= '0' && base[i] <= '9') {
		} else {
			return "NV"
		}
	}
	for _, r := range base {
		if seen[r] { // Check if the rune has been seen before
			return "NV" // Duplicate found
		}
		seen[r] = true // Mark the rune as seen
	}
	var aa int
	var bb int
	for {
		aa += n % baseLen
		n /= baseLen
		bb = aa % baseLen
		aa /= baseLen
		res = string(base[bb]) + res
		if n == 0 {
			break
		}
	}
	return neg + res
}
