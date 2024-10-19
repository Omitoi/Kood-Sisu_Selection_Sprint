package sprint

/*
Str Concat With

Learning Objectives:
*	Develop the skill to aggregate or combine data from different sources or elements, such as joining multiple strings into one.
*	Gain proficiency in manipulating strings to create meaningful compositions, such as combining and separating them based on a specific pattern or delimiter.
*	Learn how to effectively handle input parameters, such as slices and separator strings, and use them to produce desired output.
*	Practice working with delimiters or separators to format data or text in a specific way, enhancing your ability to manage data structures and text output.

Instructions
Create a Go function that takes a slice of strings and a separator string as its parameters.
The function should return a new string by concatenating all the strings in the slice, with each string separated by the provided separator.
*/

func StrConcatWith(strs []string, sep string) string {
	res := ""
	for i := 0; i < len(strs); i++ {
		if i < len(strs)-1 {
			res += strs[i] + sep
		} else {
			res += strs[i]
		} //didnt check
	}
	return res
}
