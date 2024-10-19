package main

func ExtraEnc(input string, key string) string {
	var res string
	var numKey int
	var step int
	//Generates a key of up to
	for i := 0; i < len(key); i++ {
		if int(key[i]) < 32 || int(key[i]) > 127 { //excluding characters out of scope
			return ""
		}
		numKey *= 100
		numKey += int(rune(key[i]))
	}
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 127 { //excluding characters out of scope
			return ""
		} else {
			numKey %= 93 //get a step to move character by
			step = int(rune(input[i]) + rune(numKey))
			if step > '~' {
				step -= 94
			}
			res += string(rune(step))
		}
	}
	return res
}

func ExtraEncRev(input string, key string) string {
	var res string
	var numKey int
	var step int
	//Generates a key of up to
	for i := 0; i < len(key); i++ {
		if int(key[i]) < 32 || int(key[i]) > 127 { //excluding characters out of scope
			return ""
		}
		numKey *= 100
		numKey += int(rune(key[i]))
	}
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 127 { //excluding characters out of scope
			return ""
		}
		numKey %= 93 //get a step to move character by
		step = int(rune(input[i]) - rune(numKey))
		if step < ' ' {
			step += 94
		}
		res += string(rune(step))
	}
	return res
}
