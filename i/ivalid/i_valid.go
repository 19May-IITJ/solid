package ivalid

import "math"

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return math.Pow(s.Length, 2)
}

type Cube struct {
	Square
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Length, 3)
}

type Figure interface {
	Area() float64
}

type Object interface {
	Figure
	Volume() float64
}

func AreaTotal(Figures ...Figure) float64 {
	var sum float64
	for _, s := range Figures {
		sum += s.Area()
	}
	return sum
}

func AreaVolumeTotal(Figures ...Object) float64 {
	var sum float64
	for _, s := range Figures {
		sum += s.Area() + s.Volume()
	}
	return sum
}
