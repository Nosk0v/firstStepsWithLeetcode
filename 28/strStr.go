package main

import "fmt"

func strStr(haystack string, needle string) int {
	result := -1
	var stack []byte
	for i := range haystack {
		stack = append(stack, haystack[i])
		if len(stack) >= len(needle) && string(stack[len(stack)-len(needle):]) == needle {
			result = len(stack) - len(needle)
			return result
		}
	}

	return result
}

func main() {
	haystack := "sadbutsad"
	needle := "sad"
	answer := strStr(haystack, needle)
	fmt.Println(answer)
}
