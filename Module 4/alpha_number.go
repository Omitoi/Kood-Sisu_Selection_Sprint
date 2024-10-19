package sprint

/*
Alphanumber

Learning Objectives:
*	Develop skills in transforming data between different types, such as converting integers to strings and mapping digits to letters.
*	Gain expertise in working with strings, including replacing characters, preserving signs, and formatting strings based on specific rules.
*	Practice using conditional statements to handle different cases, such as preserving the minus sign for negative numbers.
*	Learn how to map numerical values to their corresponding alphabetic representations, which is a useful skill for various types of data mapping.

Instructions
Create a Go function that takes an integer as input and returns a string representing that integer.
If the number is negative, preserve the minus sign.
Replace the digits with lowercase letters, where 0 becomes a, 1 becomes b, and so on up to 9, which becomes j.
*/

func AlphaNumber(n int) string {
	var res string
	var sign string
	if n < 0 {
		sign = "-"
		n = -n
	} else if n == 0 {
		return "a"
	}
	/*
		reversing the order of numbers in integer
		var m int
		num := n
		for num > 0 {
			m = num % 10
			rev = rev*10 + m
			num /= 10
		}
	*/

	for i := n; i > 0; i /= 10 {
		res = string(rune((i%10)+'a')) + res
	}
	return sign + res
}
