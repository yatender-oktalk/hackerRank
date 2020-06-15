package main

import "fmt"

type circle struct {
	radius float32
}

type square struct {
	length float32
}

type rectangle struct {
	width float32
	length float32
}

type areable interface {
	area() float32
}


func (c circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

func (s *square) area() float32 {
	return s.length * s.length
}

func (s *rectangle) area() float32 {
	return s.width * s.length
}

func calculateArea(a areable) {
	area := a.area()
	fmt.Println(area)
}

func main() {
	cir := circle{3.5}
	calculateArea(cir)
	
	sq := &square{5.0}
	calculateArea(sq)

	rect := &rectangle{5.0, 10.0}
	calculateArea(rect)
}
