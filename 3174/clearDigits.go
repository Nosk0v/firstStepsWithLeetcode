package main

import "fmt"

func clearDigits(s string) string {
	var result []byte

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {

			if len(result) > 0 {
				result = result[:len(result)-1]

			}

		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}

func main() {
	s := "ab1"
	result := clearDigits(s)
	fmt.Println("Текст без цифр", result)
}
