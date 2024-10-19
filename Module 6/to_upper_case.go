package sprint

/*
TO UPPER CASE

Learning Objectives:
*	Practice manipulating strings to achieve desired changes, enhancing your ability to work with text data.

Instructions
Create a Go function that takes a string as its parameter and returns a new string with each letter capitalized.
*/

func ToUpperCase(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] < 97 || s[i] > 122 {
			res += string(s[i])
		} else {
			res += string(s[i] - 32)
		}

	}
	return res
}
