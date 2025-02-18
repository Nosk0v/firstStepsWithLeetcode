package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return 1 * x
	}

	left := 0
	right := x

	for left <= right {
		middle := (left + right) / 2
		if middle*middle == x {
			return middle
		} else {
			if middle*middle > x {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return left - 1
}

func main() {
	x := 20
	answer := mySqrt(x)
	fmt.Println(answer)
}
