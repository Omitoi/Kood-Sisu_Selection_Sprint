package sprint

/*
Nth Rune

Learning Objectives:
*	Gain proficiency in transforming data between different types.
*	Learn how to manipulate strings to extract specific characters or substrings, enhancing your ability to work with text data.

Instructions
Create a Go function that takes a non-empty string as its first argument and an index as its second argument.
The function should return the rune at the specified index in the string.
*/

func NRune(s string, i int) rune {
	return rune(s[i])
}
