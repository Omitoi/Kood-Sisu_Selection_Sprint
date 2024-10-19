package sprint

/*
Square Root

Learning Objectives:
*	Gain proficiency in implementing functions that perform mathematical calculations.
*	Understand the concept of integer square roots.
*	Learn to handle specific cases, such as square roots that are not whole numbers.

Instructions
You're tasked with creating a function that takes an integer as a parameter and returns its square root if the square root is a whole number;
otherwise, it should return 0.
*/

func Wingman(n, i int) int {
	if i*i > n {
		return 0
	}
	if i*i == n {
		return i
	}
	return Wingman(n, i+1)
}

func Sqrt(n int) int {
	if n < 0 {
		return 0
	}
	return Wingman(n, 1)
}
