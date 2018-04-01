package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var iterations int
	fmt.Scanf("%d", &iterations)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < iterations; i++ {
		var inputLength int
		fmt.Scanf("%d", &inputLength)

		if scanner.Scan() {
			inputStr := scanner.Text()
			input := convertToArrOfInt(inputStr, inputLength)
			fmt.Println(getCollisions(input))
		}
	}
}

func convertToArrOfInt(input string, length int) []int {
	array := make([]int, length)
	arrayOfStr := strings.Split(input, " ")

	for i, e := range arrayOfStr {
		array[i], _ = strconv.Atoi(e)
	}

	return array
}

func getCollisions(input []int) int {
	hash := make(map[int]int, len(input))

	for _, element := range input {
		if val, ok := hash[element%10]; ok {
			hash[element%10] = val + 1
		} else {
			hash[element%10] = 0
		}
	}

	collisions := 0
	for _, value := range hash {
		collisions += value
	}

	return collisions
}
