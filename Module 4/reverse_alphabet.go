package sprint

/*
Reverse Alphabet

Learning Objectives:
*	Learn to use loops or iteration constructs to generate sequences of data and adapt them based on input values.
*	Practice using conditional statements to handle cases where the input step value is zero or negative, ensuring that the function handles these scenarios correctly.
*	Understand how to design algorithms that manipulate and transform data, considering the problem's requirements and constraints.

Instructions
Create a function that takes a step value n as input.
Starting from z in the Latin alphabet, it should return the lowercase alphabet in reverse order as a string.
For each letter, you skip n-1 letters. If n is 0 or negative, use a default step of 1.
*/
func ReverseAlphabet(step int) string {
	var res string
	if step < 0 {
		step = 1
	}
	for i := 122; i >= 97; i -= step {
		res += string(rune(i))
	}
	return res
}
