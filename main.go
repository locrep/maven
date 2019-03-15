package main

import (
	"github.com/locrep/locrep-go/person"
)

func main() {
	user := person.NewPerson()
	user.String(24)
}
