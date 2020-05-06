package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Println("Input total number count")
	scan(&N)
	s := make([]int, N)
	for i := range s {
		scan(&s[i])
	}
	printList(insertionSort(s))
}

func printList(numList []int) {
	for _, val := range numList {
		fmt.Print(" ", val)
	}
}

func insertionSort(numList []int) []int {
	for i := 1; i < len(numList); i++ {
		key := numList[i]
		j := i - 1

		for j >= 0 && numList[j] > key {
			numList[j+1] = numList[j]
			j = j - 1
		}
		numList[j+1] = key
	}
	return numList
}

func scan(a ...interface{}) {
	if _, err := fmt.Scan(a...); err != nil {
		panic(err)
	}
}
