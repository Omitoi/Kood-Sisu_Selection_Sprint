package sprint

/*
Least Common Multiple

Learning Objectives:
*	You'll likely use conditional statements and loops to implement algorithms for finding the LCM, helping reinforce your understanding of conditional logic.
*	Finding the least common multiple is a problem-solving task that encourages algorithmic thinking and an understanding of numerical operations, which are valuable in software development.
*	You may explore and implement various algorithms to find the LCM, which can teach you about algorithmic efficiency and optimization.
*	Creating a function to find the LCM promotes modular programming, a practice that enhances code organization and reusability.

Instructions
Write a Go function that takes two integers as input and returns their least common multiple (LCM). The LCM is the smallest positive integer that is divisible by both of the input integers.
*/

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func LCM(a, b int) int {
	return abs(a*b) / GCD(a, b)
}
