package main

func main() {

}

func gdcOfStrings(str1 string, str2 string) string {
	sum1 := str1 + str2
	sum2 := str2 + str1
	var output string

	if sum1 == sum2 {
		if len(str1) > len(str2) {
			gdc := countGDC(len(str1), len(str2))
			output = str1[:gdc]
		} else {
			gdc := countGDC(len(str2), len(str1))
			output = str2[:gdc]
		}
	}

	return output
}

func countGDC(num1, num2 int) int {
	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	return num1
}
