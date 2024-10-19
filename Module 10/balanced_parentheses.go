package sprint

/*
Balanced Parentheses

Learning Objectives:
*	This task provides experience in working with strings and scanning them for specific characters, which is essential for text processing and parsing.
*	Implementing a stack-like approach to check for balanced parentheses helps you understand and apply data structures in solving real-world problems.
*	Determining whether parentheses are balanced or not is a problem-solving task that involves pattern recognition and logical thinking, skills valuable in software development.

Instructions
Write a Go function that takes a string containing various characters, including parentheses ()[]{}, and checks if the parentheses are balanced. 
The function should return a boolean value indicating whether the parentheses are balanced or not.
*/

func BalancedParentheses(input string) bool {
	// Stack to store opening parentheses
	stack := []rune{}

	// Map to match closing parentheses with their corresponding opening parentheses
	matchingParen := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// Iterate through each character in the input string
	for _, char := range input {
		// If it's an opening parenthesis, push it onto the stack
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		}

		// If it's a closing parenthesis, check for balance
		if char == ')' || char == ']' || char == '}' {
			// If stack is empty or the top of the stack doesn't match the closing parenthesis
			if len(stack) == 0 || stack[len(stack)-1] != matchingParen[char] {
				return false
			}
			// Pop the top of the stack
			stack = stack[:len(stack)-1]
		}
	}

	// If the stack is empty, all parentheses were matched; otherwise, it's unbalanced
	return len(stack) == 0
}
