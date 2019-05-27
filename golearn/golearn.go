package main

import (
	"fmt"
)

func main() {
	fmt.Println("I am go function")
	x := 2
	x1 := addByMem(&x)
	fmt.Println("x -> ", x)
	fmt.Println("x1 -> ", x1)
}

func addByMem(a *int) int {
	*a = *a + 1
	return *a
}
