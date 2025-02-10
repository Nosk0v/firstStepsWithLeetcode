package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		if idx, found := hmap[diff]; found {
			return []int{idx, i}
		}
		hmap[num] = i
	}
	return nil
}

func main() {
	nums := []int{1, 2, 3, 4, 6}
	target := 7
	result := twoSum(nums, target)

	if result != nil {
		fmt.Println("Индексы", result)
	} else {
		fmt.Println("Пара не найдена")
	}

}
