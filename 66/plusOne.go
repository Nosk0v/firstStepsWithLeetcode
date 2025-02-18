package main

import "fmt"

func plusOne(digits []int) []int {

	lenS := len(digits) - 1
	for i := lenS; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits
}

func main() {
	digits := []int{1, 2, 9}
	answer := plusOne(digits)
	fmt.Println(answer)
}
