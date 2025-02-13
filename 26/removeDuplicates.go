package main

import "fmt"

func removeDuplicates(nums []int) int {
	result := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[result] = nums[i]
			result++
		}
	}

	return result
}

func main() {
	nums := []int{1, 1, 2}
	answer := removeDuplicates(nums)
	fmt.Println(answer)
}
