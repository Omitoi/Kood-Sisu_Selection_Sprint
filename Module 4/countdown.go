package sprint

/*
Countdown

Learning Objectives:
*	Develop the skill of using loops, like for loops, to iterate through sequences of data, providing control over how data is processed and manipulated.
*	Gain expertise in working with strings in Go, including techniques to concatenate, modify, and manipulate strings effectively to achieve the desired outcomes.
*	Practice using conditional statements, like if statements, to make decisions in your code.
	Understand how to control the flow of your program based on specific conditions or criteria.

Instructions
Create a function that takes a one-digit integer as input and returns a countdown string.
The countdown should start from the given number, skip every second number, and end with 0!.
For example, if the input is 7, the function should return "7, 5, 3, 1, 0!".
*/

func Countdown(n int) string {
	var res string
	for i := n; i > 0; i -= 2 {
		res += string(rune(i+48)) + ", "
	}
	return res + "0!"
}
