package main

import (
	"fmt"
	"hackerRank/hackerRank/libraries"
)

// Person struct to use at other places
type Person struct {
	Name    string
	Age     int
	Surname string
	Hobbies []string
	id      string
}

// Dog type struct
type Dog struct {
	Name    string
	Age     int
	Surname string
	Hobbies []string
	id      string
}

// GetFullName is person get full name from this
func (person *Person) GetFullName() string {
	return person.Name + " " + person.Surname
}

// GetFullName is dog get full name from this
func (dog *Dog) GetFullName() string {
	return dog.Name
}

func main() {
	addFive := adder(5)
	addSix := adder(6)

	k := addFive(9)
	j := addSix(9)

	fmt.Println(k)
	fmt.Println(j)

	p := Person{
		Name:    "Mario",
		Age:     23,
		Surname: "Castro",
		Hobbies: []string{"cycling"},
		id:      "abc-123",
	}

	d := Dog{
		Name:    "DMario",
		Age:     15,
		Surname: "DCastro",
		Hobbies: []string{"barking"},
		id:      "abc-123",
	}

	fmt.Printf("%s likes %s\n", p.GetFullName(), p.Hobbies[0])
	fmt.Printf("%s likes %s", d.GetFullName(), d.Hobbies[0])

	sumRes := arithmetic.Sum(5, 6)
}

func adder(i int) func(int) int {
	return func(k int) int {
		return i + k
	}
}
