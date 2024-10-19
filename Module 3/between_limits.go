package sprint

/*
Between Limits

Learning Objectives:
*	Develop a solid understanding of working with runes in Go, including comparing, extracting, and modifying them effectively.
*	Gain expertise in building and manipulating strings in Go, with the ability to construct strings containing a range of runes to meet specific requirements.
*	Practice processing input runes in a way that accommodates different orders or variations, ensuring that the function works correctly for various scenarios.
*	Gain proficiency in using loops or iterations to generate sequences of runes between two input runes, demonstrating control over the process.

Instructions
Build a function that accepts two runes as input and returns a string containing all the runes that come between these two runes in the alphabet, in the correct order.
For instance, if the input runes are 'f' and 'j', the function should return "ghi" regardless of the order in which the runes are given.
*/

func BetweenLimits(from, to rune) string {
	var res string
	var first int
	var second int
	input1 := int(from)
	input2 := int(to)
	if input1 > input2 {
		first = input2
		second = input1
	} else if input1 < input2 {
		first = input1
		second = input2
	} else {
		return ""
	}

	first++
	for i := first; i < second; i++ {
		val := rune(i)
		letter := string(val)
		res += letter
	}
	return res
}
