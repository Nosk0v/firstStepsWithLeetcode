package main

import "fmt"

func isPalindrome(x int) bool {
	if x >= 0 && x < 10 {
		return true
	}

	if x < 0 || x%10 == 0 {
		return false
	}
	original := x
	reverse := 0

	for x > 0 {
		digit := x % 10
		reverse = reverse*10 + digit
		x /= 10
	}

	return reverse == original
}
func main() {
	x := 23
	result := isPalindrome(x)
	fmt.Println(result)
}
