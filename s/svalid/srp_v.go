package svalid

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

/*
A type or a function must have only one & one
job & one & one responsibilty but code in this file voilates this
*/

type Circle struct {
	radius float32
}

func (c Circle) Name() string {
	return "circle"
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

type Figure struct {
	length float32
}

func (s Figure) Name() string {
	return "Figure"
}

func (s Figure) Area() float32 {
	return s.length * s.length
}

type FiguresI interface {
	Area() float32
	Name() string
}

type Printer struct {
}

func (p Printer) Text(s FiguresI) string {
	return fmt.Sprintf("area  of the %s: %f", s.Name(), s.Area())
}

func (p Printer) JSON(s FiguresI) string {
	response := struct {
		Name string  `json:"figure"`
		Area float32 `json:"area"`
	}{
		Name: s.Name(),
		Area: s.Area(),
	}

	serialized, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	return string(serialized)
}
