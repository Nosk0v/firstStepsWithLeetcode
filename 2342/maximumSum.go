package main

import (
	"fmt"
	"math"
)

func maximumSum(nums []int) int {
	answer := -1
	hmap := make(map[int]int)
	for _, num := range nums {
		sum := getSumDigits(num)

		if value, found := hmap[sum]; found {
			answer = int(math.Max(float64(answer), float64(value+num)))
		}
		hmap[sum] = int(math.Max(float64(hmap[sum]), float64(num)))
	}
	return answer
}

func getSumDigits(num int) int {
	sumDigits := 0
	for num > 0 {
		sumDigits += num % 10
		num /= 10
	}
	return sumDigits
}
func main() {
	var nums []int
	nums = append(nums, 18, 43, 36, 13, 7)
	result := maximumSum(nums)
	fmt.Println(result)
}
