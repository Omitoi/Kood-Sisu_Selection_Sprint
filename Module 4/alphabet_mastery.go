package sprint

/*
Alphabet Mastery

Learning Objectives:
*	Learn how to use loops or iteration to generate a sequence of letters, adapting the loop based on the input value.
*	Gain proficiency in working with strings to construct a string of letters, using indexing and other string manipulation techniques.

Instructions
Create a function that takes a positive integer as input and returns the corresponding number of letters from the Latin alphabet.
The input integer won't be larger than the alphabet's length.
*/

func AlphabetMastery(n int) string {
	var res string
	for i := 97; i < n+97; i++ {
		res += string(rune(i))
	}
	return res
}
