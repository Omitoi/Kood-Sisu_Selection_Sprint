package sprint

/*
To Capital Case

Learning Objectives:
*	Practice manipulating strings to achieve desired changes, enhancing your ability to work with text data.

Instructions
Create a Go function that takes a string as its parameter.
The function should capitalize the first letter of each word while converting the rest of the word to lowercase.
In this task a word is defined as a sequence of alphanumeric characters.
*/

func ToCapitalCase(s string) string {
	res := ""
	cap := true

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			if cap {
				res += string(s[i] - 32)
				cap = false
			} else {
				res += string(s[i])
			}
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			if !cap {
				res += string(s[i] + 32)
			} else {
				res += string(s[i])
				cap = false
			}
		} else if s[i] >= '0' && s[i] <= '9' {
			if !cap {
				res += string(s[i])
			} else {
				res += string(s[i])
				cap = false
			}
		} else {
			res += string(s[i])
			cap = true
		}
	}
	return res
}
