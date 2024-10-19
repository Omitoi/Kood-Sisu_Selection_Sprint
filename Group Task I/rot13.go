package main

func Rot13(input string) string {
	var res string
	for i := 0; i < len(input); i++ {
		switch {
		case input[i] >= 'A' && input[i] <= 'Z':
			res += string('A' + (input[i]-'A'+13)%26)
		case input[i] >= 'a' && input[i] <= 'z':
			res += string('a' + (input[i]-'a'+13)%26)
		case input[i] >= '0' && input[i] <= '9':
			res += string('0' + (input[i]-'0'+5)%10) // fixed numbers for rot+5
		case input[i] == ' ':
			res += string(' ') // fixed space
		default:
			res += string(input[i]) // Include non-alphabetic characters HOPEFULLY
		}
	}
	return res
}

func Rot13Rev(input string) string {
	var res string
	for i := 0; i < len(input); i++ {
		switch {
		case input[i] >= 'A' && input[i] <= 'Z':
			res += string('A' + (input[i]-'A'+13)%26)
		case input[i] >= 'a' && input[i] <= 'z':
			res += string('a' + (input[i]-'a'+13)%26)
		case input[i] >= '0' && input[i] <= '9':
			res += string('0' + (input[i]-'0'+5)%10) // fixed numbers for rot+5
		case input[i] == ' ':
			res += string(' ') // fixed space
		default:
			res += string(input[i]) // Include non-alphabetic characters HOPEFULLY
		}
	}
	return res
}
