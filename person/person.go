package person

import (
	"fmt"
	"strconv"
)

type Person struct {
	name    string
	surname string
	age     int
}

func NewPerson(age int) Person {
	return Person{
		"ahmet",
		"dalgic",
		age,
	}
}

func (p Person) Age() int {
	return p.age
}

func (p Person) String(age int) {
	fmt.Println(p.name + " " + p.surname + strconv.Itoa(age))
}
