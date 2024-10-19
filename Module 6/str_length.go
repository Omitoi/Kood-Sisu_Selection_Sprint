package sprint

/*
Str Length

Learning Objectives:
*	Develop the ability to work with strings, including counting runes and bytes.
*	Gain an understanding of character encoding schemes, with a specific focus on UTF-8 encoding and how it impacts the length of strings.

Instructions
Create a Go function that accepts a string and returns two integers.
The first integer represents the number of runes in the string, and the second integer represents the byte length of the string.
The function should handle UTF-8 encoding for all characters.
*/

func StrLength(s string) []int {
	runeCount := 0
	for range s {
		runeCount++
	}
	return []int{runeCount, len(s)}
}
