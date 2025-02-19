package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	var strings []byte
	lenS := len(s) - 1
	for i := lenS; i >= 0; i-- {

		if s[i] == ' ' && len(strings) > 0 {
			return len(strings)
		}
		if s[i] == ' ' {
			continue
		} else {
			strings = append(strings, s[i])
		}

	}
	return len(strings)
}

func main() {
	s := "a "
	result := lengthOfLastWord(s)
	fmt.Println("Длина слова", result)
}
