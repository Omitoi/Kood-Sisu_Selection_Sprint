package sprint

/*
Next Prime

Learning Objectives:
*	Learn optimization techniques to avoid time-outs for large numbers.
*	Understand the concept of prime numbers and how to find them efficiently.

Instructions
You're tasked with creating a function that takes an integer as a parameter. If the given integer is a prime, the function should return that integer; otherwise, it should find and return the next prime number.
You should optimize the function to avoid time-outs during testing.
Remember that prime numbers can be only natural numbers greater than 1.
*/

func NextPrime(n int) int {
	if n == 0 {
		return 2
	}
	if n%2 == 0 {
		return NextPrime(n + 1)
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return NextPrime(n + 1)
		}
	}
	return n
}
