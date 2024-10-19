package sprint

/*
Get First Rune

Learning Objectives:
*	Gain proficiency in transforming data between different types.
*	Learn how to manipulate strings to extract specific characters or substrings, enhancing your ability to work with text data.

Instructions
Create a Go function that takes a non-empty string as input and returns the first character of the string as a rune.
*/

func GetFirstRune(s string) rune {
	return rune(s[0])
}
