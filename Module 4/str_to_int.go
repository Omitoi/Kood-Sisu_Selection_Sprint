package sprint

/*
String to Integer

Learning Objectives:
*	Develop the skill of converting data between different types, such as converting strings to integers, handling signs, and recognizing invalid input.
*	Understand how to handle potential errors or non-valid input gracefully by returning a default value (0 in this case) instead of raising an error.
*	Practice using conditional statements to make decisions in your code, especially when handling signs and different cases.
*	Learn techniques for parsing and interpreting strings to extract meaningful information, such as numbers from mixed character strings.

Instructions
Create a Go function that mimics the behavior of the Atoi function.
Atoi converts a string representing a number into an integer.
*	The function should return the integer value.
*	If the input string is not a valid number, it should return 0.
*	The function must handle signs such as + or -.
*	The exercise does not require you to return an error.
*/

func StrToInt(s string) int {
	var neg bool
	var res int
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
			res *= 10
			res += int(rune(s[i]) - 48)
		} else if rune(s[i]) == 45 && i == 0 {
			neg = true
		} else if rune(s[i]) == 43 && i == 0 {
			neg = false
		} else {
			return 0
		}
	}
	if neg {
		res = -res
	}
	return res
}
