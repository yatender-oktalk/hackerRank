package main

import "fmt"

func main() {
	m := 8
	n := 2

	gcd := calculateGcd(m, n)
	fmt.Println(gcd)
}

func calculateGcd(m, n int) int {
	if n != 0 {
		rem := m % n
		return calculateGcd(n, rem)
	}
	return m
}
