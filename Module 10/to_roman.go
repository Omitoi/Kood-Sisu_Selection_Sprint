package sprint

/*
To Roman

Learning Objectives:
*	Implementing input validation by checking the validity of the input integer is an important aspect of this task and teaches you about handling edge cases.
*	This task involves performing mathematical computations to convert a valid integer into its Roman numeral representation.
*	You'll use conditional statements to check the input and determine how to convert specific numbers into Roman numerals.
*	Converting numbers to Roman numerals requires string manipulation, which is a fundamental skill for working with text data.
*	Converting valid integers to Roman numerals is a problem-solving task that encourages logical thinking and an understanding of number systems, valuable skills in software development.

Instructions
Write a Go function that takes an integer and converts it into a Roman numeral.
The function should return "Invalid input" if the input integer is less than 1 or more than 3999.
Otherwise, it should return the Roman numeral representation of the valid input integer.
*/

func ToRoman(num int) string {
	if num < 1 || num > 3999 {
		return "Invalid input"
	}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			res += symbols[i]
			num -= values[i]
		}
	}

	return res
}
