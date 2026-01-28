package main

import "strings"

func main() {

}

func mergeAlternately(word1 string, word2 string) string {
	var output strings.Builder
	output.Grow(len(word1) + len(word2)) // изначальная емкость strings.Builder
	i := 0
	for i < len(word1) || i < len(word2) {
		if i < len(word1) {
			output.WriteByte(word1[i])
		}
		if i < len(word2) {
			output.WriteByte(word2[i])
		}
		i++
	}
	return output.String()
}
