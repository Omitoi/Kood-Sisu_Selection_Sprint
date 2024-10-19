package sprint

/*
Are Anagrams?

Learning Objectives:
*	You'll practice working with strings, which is a fundamental skill in text processing and data manipulation.
*	Writing a function to determine if two strings are anagrams involves applying conditional logic to compare character arrangements.
*	Ignoring character case and whitespaces in the comparison of anagrams highlights the importance of handling different cases and spaces in real-world text data.
*	Detecting anagrams is a problem-solving task that encourages logical thinking and pattern recognition, which are important skills in software development.

Instructions
Write a Go function that takes two strings as input and returns a boolean value indicating whether the input strings are anagrams or not. 
Anagrams are words or phrases formed by rearranging the letters of another, and in this task, you should ignore differences in character case and whitespace.

Allowed packages
For this task you're allowed to use:
*	sort
*	strings
*/

func AreAnagrams(str1, str2 string) bool {
	seen := make(map[rune]bool)
	for _, r := range str1 {
		if r >= 97 && r <= 122 {
			r -= 32 //lowercase conversion
		}
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '\v' || r == '\f' {

		} else if r >= 'A' && r <= 'Z' {
			seen[r] = true
		}
		println(string(r), seen[r])
	}
	for _, r := range str2 {
		if r >= 97 && r <= 122 {
			r -= 32 //lowercase conversion
		}
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '\v' || r == '\f' {

		} else if r >= 'A' && r <= 'Z' {
			if seen[r] {
				seen[r] = false
			}
		}
		println(string(r), seen[r])
	}
	for _, r := range seen {
		println(r)
		if r {
			return false
		}
	}
	return true
}
