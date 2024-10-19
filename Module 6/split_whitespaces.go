package sprint

/*
Split on Whitespaces

Learning Objectives:
*	Develop the ability to parse and extract information from strings based on specific delimiters or separators.
*	Gain proficiency in working with string slices to store and manage collections of words or elements from a text.
*	Practice handling and processing different types of whitespace characters, such as spaces, tabs, and newlines, in text data.

Instructions
Create a Go function that takes a string as its parameter. The function should split the string into words and store them in a string slice. Words are separated by spaces, tabs, and newlines.
*/

func SplitWhitespaces(s string) []string {
	word := ""
	var res []string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r' || s[i] == '\v' || s[i] == '\f' {
			res = append(res, word)
			word = ""
		} else {
			word += string(s[i])
		}
	}
	res = append(res, word)
	return res
}
