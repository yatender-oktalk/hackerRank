package main

func main() {
	addSix := addN(6)
	addTen := addN(10)

	println("adding six into 7", addSix(7))
	println("adding ten into 7", addTen(7))
}

func addN(n int) func(int) int {
	return func(m int) int {
		return m + n
	}
}
