package person

import "testing"

func TestPrintPersonAge(t *testing.T){
	person := NewPerson(25)

	if person.Age() != 25{
		t.Fail()
	}
}
