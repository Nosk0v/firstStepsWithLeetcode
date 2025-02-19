package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num == target || num > target {
			return i
		}
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	return 0
}

func main() {
	nums := []int{1, 2, 3, 4}
	target := 5
	answer := searchInsert(nums, target)
	fmt.Println(answer)
}
