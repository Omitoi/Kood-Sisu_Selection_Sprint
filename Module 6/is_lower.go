package sprint

/*
Lower or Not?

Learning Objectives:
*	Develop the skill of validating data to determine if it meets specific criteria or conditions.
*	Practice using conditional statements to make decisions in your code based on data characteristics.
*	Learn how to process and analyze strings effectively, enabling you to check and manipulate the contents of text data.
*	Understand how to return boolean values from functions to indicate whether certain conditions are met or not.

Instructions
Create a Go function that takes a string as its parameter. The function should return true if the string contains only lowercase characters, and false otherwise.
*/

func IsLower(s string) bool {
	res := true
	for i := 0; i < len(s); i++ {
		if s[i] < 97 || s[i] > 122 {
			res = false
		}
	}
	return res
}
