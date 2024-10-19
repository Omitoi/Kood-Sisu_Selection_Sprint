package main

func Reverse1(message string) string {
	var reverse string
	for i := 0; i < len(message); i++ {
		if message[i] >= 32 && message[i] <= 47 || message[i] >= 58 && message[i] <= 64 || message[i] >= 91 && message[i] <= 96 || message[i] >= 123 && message[i] <= 127 {
			reverse += string(rune(message[i]))
		}
		if message[i] >= 'a' && message[i] <= 'z' {
			reverse += string('z' - (rune(message[i] - 'a')))
		}
		if message[i] >= 'A' && message[i] <= 'Z' {
			reverse += string('Z' - (rune(message[i] - 'A')))
		}
		if message[i] >= '0' && message[i] <= '9' {
			reverse += string('9' - (rune(message[i] - '0')))
		}
	}
	return reverse
}
