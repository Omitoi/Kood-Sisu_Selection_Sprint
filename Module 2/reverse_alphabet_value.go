package sprint

/*
Reverse Alphabet Value

Learning Objectives:
*	You will gain experience in manipulating characters in Go, a fundamental skill for text processing.
*	This task requires you to think creatively to achieve the desired transformation. 
	Problem-solving is a crucial skill for software development and programming.

Instructions
Write a Go function that takes a lowercase letter rune ('a'-'z') as input and returns its reverse letter in the Latin alphabet as a rune integer value.
For example, 'a' should be transformed to 'z', 'y' should become 'b', and so on. The function should maintain the case of the input letter.
*/

func ReverseAlphabetValue(ch rune) rune {
	return 122 - (ch - 97)
}
