package main

import (
	"fmt"
	"strconv"
)

func main() {
	var iter int
	fmt.Scanf("%d", &iter)

	for i := 0; i < iter; i++ {
		var upperLimit, coolLimit int
		fmt.Scanln(&upperLimit, &coolLimit)

		coolNumbers := 0
		for j := 1; j <= upperLimit; j++ {
			coolness := getCoolness(j)
			if coolness >= coolLimit {
				coolNumbers++
			}
		}

		fmt.Println(coolNumbers)
	}
}

func getCoolness(input int) int {
	n := int64(input)
	binary := strconv.FormatInt(n, 2)

	return countOccurrences(binary)
}

func countOccurrences(input string) int {
	count := 0
	for i := 0; i < len(input)-2; i++ {
		if input[i:i+3] == "101" {
			count++
		}
	}
	return count
}
