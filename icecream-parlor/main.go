package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner *bufio.Scanner

func main() {
	readInput()
}

func readInput() {
	var tripCount int
	fmt.Scanln(&tripCount)

	scanner = bufio.NewScanner(os.Stdin)
	for i := 0; i < tripCount; i++ {
		var moneyPooled int
		var flavors int
		var costs []Cost

		scanInt(&moneyPooled)
		scanInt(&flavors)
		scanCosts(&costs, flavors)

		trip := &Trip{
			MoneyPooled: moneyPooled,
			Costs:       costs,
		}

		sunnyCost, johnnyCost := trip.findOptimizedCost()

		fmt.Printf("%d %d\n", sunnyCost, johnnyCost)
	}
}

func scanInt(i *int) {
	if scanner.Scan() {
		inputStr := scanner.Text()
		*i, _ = strconv.Atoi(inputStr)
	}
}

func scanCosts(costs *[]Cost, size int) {
	if scanner.Scan() {
		inputStr := scanner.Text()
		*costs = parseStringArray(inputStr, size)
	}
}

func parseStringArray(input string, length int) []Cost {
	array := make([]Cost, length)
	arrayOfStr := strings.Split(input, " ")

	for i, e := range arrayOfStr {
		value, _ := strconv.Atoi(e)
		array[i] = Cost{index: i + 1, value: value}
	}

	return array
}

type Trip struct {
	MoneyPooled int
	Costs       []Cost
}

type Cost struct {
	index int
	value int
}

type ByCost []Cost

func (a ByCost) Len() int {
	return len(a)
}

func (a ByCost) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByCost) Less(i, j int) bool {
	return a[i].value < a[j].value
}

func (t *Trip) findOptimizedCost() (int, int) {
	sort.Sort(ByCost(t.Costs))
	costs := t.Costs
	for i, j := 0, len(costs)-1; i < j; {
		sum := costs[i].value + costs[j].value
		if sum == t.MoneyPooled {
			return costs[i].index, costs[j].index
		} else if sum > t.MoneyPooled {
			j--
		} else if sum < t.MoneyPooled {
			i++
		}
	}
	return 0, 0
}
