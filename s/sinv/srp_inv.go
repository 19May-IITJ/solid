package sinv

import (
	"fmt"
	"math"
)

/*
A type or a function must have only one & one
job & one & one responsibilty but code in this file voilates this
*/

type Circle struct {
	Radius float32
}

/*
Voilation of single responsibility principle
Calculated area and printing the same in same method call
*/
func (c *Circle) Area() {
	fmt.Printf("Area of Circle: %f\n", math.Pi*c.Radius*c.Radius)
}

type Square struct {
	Length float32
}

func (s *Square) Area() {
	fmt.Printf("Area of Square: %f\n", s.Length*s.Length)
}
