package main

import "fmt"

func removeOccurrences(s string, part string) string {
	var stack []byte
	for i := range s {
		stack = append(stack, s[i])
		if len(stack) >= len(part) && string(stack[len(stack)-len(part):]) == part {
			stack = stack[:len(stack)-len(part)]
		}
	}
	return string(stack)
}

func main() {
	s := "daabcbaabcbc"
	part := "abc"
	result := removeOccurrences(s, part)
	fmt.Println(result)
}
