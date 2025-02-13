package main

import "fmt"

func removeElement(nums []int, val int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[result] = nums[i]
			result++
		}
	}
	return result

}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 2
	answer := removeElement(nums, val)
	fmt.Println(answer)
}
