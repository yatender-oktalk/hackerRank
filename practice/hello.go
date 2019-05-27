package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var i int
	var c, python, java bool
	fmt.Println("Helo")
	fmt.Println("The time is", time.Now())
	fmt.Println("My Favorite number is ", rand.Intn(500))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(49))
	fmt.Printf("Value of Pie is %g\n", math.Pi)
	fmt.Printf("value of calling a sum function %f\n", addition(3, 5))
	var a, b int = swap(3, 5)
	fmt.Println("value of calling a swap", a, b)
	var x, y int = addOne(3, 5)
	fmt.Println(x, y)
	fmt.Println(i, c, python, java)

	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4, 5)
	printSlice(s)

	var pow = []int{1, 2, 3, 4, 5, 5, 6}
	iterateNicely(pow)
}

func iterateNicely(itemArray []int) {
	for i, v := range itemArray {
		fmt.Printf("2**%d = %d \n", i, v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}

func addition(a, b int) float32 {
	return float32(a + b)
}

func swap(a, b int) (int, int) {
	return b, a
}

func addOne(a, b int) (x, y int) {
	x = a + 1
	y = b + 1
	return
}
