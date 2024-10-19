package sprint

/*
Get Last Rune

Learning Objectives:
*	Gain proficiency in transforming data between different types.
*	Learn how to manipulate strings to extract specific characters or substrings, enhancing your ability to work with text data.

Instructions
Create a Go function that takes a non-empty string as input and returns the last character of the string as a rune.
*/

func GetLastRune(s string) rune {
	return rune(s[len(s)-1])
}
