package main

import "fmt"

type gender int

const (
	male gender = iota
	female
)

type person struct {
	name   string
	age    int
	gender gender
}

func (p person) print() {
	fmt.Printf("%s/%v (%d)\n", p.name, p.gender, p.age)
}

func (p employee) print() {
	fmt.Println("employee info:")
	p.person.print()
}

// struct embedding
type employee struct {
	person
	department string
}

func main() {
	println()
	scryner := employee{
		person: person{
			name:   "scryner",
			age:    38,
			gender: male,
		},
		department: "Naver Labs",
	}

	scryner.print()
}
