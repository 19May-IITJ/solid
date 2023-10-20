package lvalid

import "fmt"

/*
this talks about abstract interface defination and its re-usability
*/
type Homosapien struct {
	Name string
}

func (h Homosapien) GetName() string {
	return h.Name
}

type Teacher struct {
	Homosapien
	Qualification string
	Salary        float64
}

type Student struct {
	Homosapien
	Marks map[string]int
}

type Person interface {
	GetName() string
}

type Printer struct {
}

func (Printer) Info(p Person) {
	fmt.Println("Name: ", p.GetName())
}
