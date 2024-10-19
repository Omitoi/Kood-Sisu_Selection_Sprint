package sprint

/*
String Compressor

Learning Objectives:
*	This task will provide practical experience in working with strings, which is essential for various text processing tasks.
*	You will use iteration and counting to detect and compress consecutive repeating characters while correctly handling the edge case of single characters.
*	Building a compressed string from the counts and characters is a useful skill for text transformation.
*	Implementing string compression involves problem-solving skills, such as identifying patterns and devising efficient solutions while considering edge cases.

Instructions
Write a Go function that takes a string as input and compresses it by replacing consecutive repeating characters with a count of repetitions followed by the character. However, if there is only one character in a sequence, it should not be prepended with a number. For example, "aaabbaac" should be compressed to "3a2b2ac", but "abcde" should remain unchanged as "abcde".

Allowed packages
For this task you're allowed to use:
fmt
*/

import "fmt"

func StrCompress(input string) string {
	res := ""
	if len(input) <= 1 {
		return input
	}
	counter := 1
	for j := 0; j < len(input)-1; j++ {
		if input[j] == input[j+1] {
			counter++
		} else {
			if counter > 1 {
				res += fmt.Sprintf("%d%v", counter, string(input[j]))
			} else {
				res += string(input[j])
			}
			counter = 1
		}
	}
	if input[len(input)-1] == input[len(input)-2] {
		res += fmt.Sprintf("%d%v", counter, string(input[len(input)-1]))
	} else {
		res += string(input[len(input)-1])
	}
	return res
}
