package sprint

/*
Shift By

Learning Objectives:

*	Understand rune data type, and rune literal.
*	Learn how to work with runes in Go, including extracting, modifying, and returning runes to achieve the desired shift.
*	Develop the ability to create a well-structured Go function that takes input parameters, performs operations on them, and returns the desired result.
*	Practice handling edge cases, such as ensuring that the function works correctly with lowercase letters and correctly loops back to the start of the alphabet if the step value exceeds 26.

Instructions
Create a Go function that takes a single lowercase letter as a rune and an int 'step'.
The function should shift the letter in the alphabet by the specified 'step' value, and return the resulting letter.
For example, if you shift 'a' by 4 steps, it should become 'e'. Remember to handle looping back to the start of the alphabet if needed.
Only lower case characters are tested.
*/

func ShiftBy(r rune, step int) string {

	step %= 26                 //reducing to a range 1-26 to stay within the range of lowercase letters
	nextrune := r + rune(step) //adding the step

	//handling overflow
	if nextrune >= 123 {
		nextrune -= 26
	}

	return string(nextrune)
}
