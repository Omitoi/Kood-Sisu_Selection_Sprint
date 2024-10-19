package sprint

/*
From Roman

Learning Objectives:
*	This task involves working with strings and performing operations to convert a Roman numeral into an integer.
*	You'll need to perform mathematical computations to interpret the Roman numeral and convert it into an integer.
*	Converting Roman numerals to integers may involve conditional statements to handle different cases and combinations of Roman numerals.
*	Converting Roman numerals to integers is a problem-solving task that encourages logical thinking and understanding of number systems, valuable skills in software development.

Instructions
Write a Go function that takes a valid Roman numeral as input and converts it into an integer. The function should return the integer representation of the valid Roman numeral input.
*/

func FromRoman(s string) int {
	romanValues := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	res := 0
	length := len(s)
	for i := 0; i < length; i++ {
		currentValue := romanValues[rune(s[i])]
		if i+1 < length && currentValue < romanValues[rune(s[i+1])] {
			res -= currentValue
		} else {
			res += currentValue
		}
	}

	return res
}
