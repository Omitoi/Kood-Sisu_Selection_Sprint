package sprint

/*
Str Split By

Learning Objectives:
*	Develop the ability to parse and extract information from strings based on specific delimiters or separators.
*	Gain proficiency in working with string slices to store and manage collections of words or elements from a text.
*	Practice handling and processing different types of delimiters, such as characters or strings, used to separate parts of a text.

Instructions
Create a Go function that takes a string s and a separator sep as parameters. 
The function should split the string s into substrings using the separator sep and return a slice of strings containing the resulting substrings.
*/

func StrSplitBy(s, sep string) []string {
	var res []string
	var word string
	sepLen := len(sep)

	for i := 0; i < len(s); {
		if i <= len(s)-sepLen && s[i:i+sepLen] == sep {
			res = append(res, word)
			word = ""
			i += sepLen
		} else {
			word += string(s[i])
			i++
		}
	}
	if word != "" {
		res = append(res, word)
	}
	if res == nil {
		return []string{}
	}
	return res
}
