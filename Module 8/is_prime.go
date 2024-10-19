package sprint

/*
Prime Numbers

Learning Objectives:
*	Learn optimization techniques to avoid time-outs for large numbers.
*	Understand the concept of prime numbers and their characteristics.

Instructions
You're tasked with creating a function that takes an integer as a parameter and returns true if this number is prime and false if it's not.
You should optimize the function to avoid time-outs during testing.
Remember that prime numbers can be only natural numbers greater than 1.
*/

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
