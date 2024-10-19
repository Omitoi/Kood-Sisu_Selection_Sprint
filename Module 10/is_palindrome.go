package sprint

/*
Is Palindrome?

Learning Objectives:
*	You'll practice working with strings, which is a fundamental skill in text processing and data manipulation.
*	Writing a function to determine if a string is a palindrome or not involves applying conditional logic to compare characters in reverse order.
*	The task includes not changing the character case, which teaches you to pay attention to case sensitivity in programming.

Instructions
Write a Go function that takes a string as input and returns a boolean value indicating whether the input string is a palindrome or not. 
A palindrome is a string that reads the same forwards and backward, considering character case and white spaces.
*/

func IsPalindrome(s string) bool {
	sLen := len(s) - 1
	for i := 0; i < len(s)-1; i++ {
		for j := sLen; j > 0; j-- {
			if s[i] != s[j] {
				return false
			}
			i++
		}
	}
	return true
}
