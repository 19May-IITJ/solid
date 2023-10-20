package ovalid

import "math"

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct {
	Length float32
}

func (s Square) Area() float32 {
	return s.Length * s.Length
}

type Triangle struct {
	Height float32
	Base   float32
}

func (t Triangle) Area() float32 {
	return t.Base * t.Height / 2
}

type FigureI interface {
	Area() float32
}

type Processor struct {
}

func (p Processor) AreaSum(Figures ...FigureI) float32 {
	var sum float32
	for _, Figure := range Figures {
		sum += Figure.Area()
	}
	return sum
}
