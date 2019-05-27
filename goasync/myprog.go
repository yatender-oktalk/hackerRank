package main

import (
	"fmt"
)

var slice = []int{1, 2, 3, 4, 5}

type testInt func(int) bool

func isEven(item int) bool {
	return item%2 == 0
}

func isOdd(item int) bool {
	return (item % 2) != 0
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	fmt.Println(filter(slice, isOdd))
	fmt.Println("I am starting the prog")
	fmt.Println(filter(slice, isEven))
}
