package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width  float64
	length float64
}

type circle struct {
	radius float64
}

type person struct {
	name string
	age  int
}

func (r Rectangle) area() float64 {
	return r.width * r.length
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	p := person{age: 22, name: "Yatender"}
	fmt.Println(p.name, p.age)
	rect := Rectangle{5, 6}
	rectArea := rect.area()
	fmt.Println(rectArea)
	cir := circle{5}
	fmt.Println(cir.area())
}
