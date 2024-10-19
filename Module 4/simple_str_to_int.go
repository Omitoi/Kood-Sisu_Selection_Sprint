package sprint

/*
Simple String to Integer

Learning Objectives:
*	Develop the ability to convert data between different types, such as converting strings to integers or vice versa.
*	Understand how to validate input data and handle potential errors or invalid data gracefully.
*	Gain proficiency in working with strings, including extracting, modifying, or filtering characters as needed.
*	Learn how to design algorithms for specific tasks, considering the problem`s requirements and constraints.
Instructions
Create a Go function that converts a valid string representation of a number into an integer.
If the string is not a valid number, the function should return 0.
The input will only contain one or more digits, and you don't need to handle signs like + or -.
*/

func SimpleStrToInt(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res *= 10
		res += int(rune(s[i]) - 48)
	}
	return res
}
