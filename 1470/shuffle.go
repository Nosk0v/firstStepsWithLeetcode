package main

import "fmt"

func shuffle(nums []int, n int) []int {
	var stack []int

	for i, num := range nums {
		if i == n {
			return stack
		}
		stack = append(stack, num)
		stack = append(stack, nums[i+n])
	}
	return stack
}

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	answer := shuffle(nums, n)
	fmt.Println(answer)
}
