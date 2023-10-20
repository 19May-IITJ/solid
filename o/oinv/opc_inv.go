package oinv

import "math"

/*
Open for extension and closed for modification
*/

type Circle struct {
	Radius float64
}

type Square struct {
	Length float64
}

type Processor struct {
}

func (p Processor) AreaSum(shapes ...interface{}) float64 {
	var sum float64
	for _, shape := range shapes {
		switch types := shape.(type) {
		case Circle:
			r := types.Radius
			sum += math.Pi * r * r
		case Square:
			l := types.Length
			sum += l * l
		}
	}
	return sum
}
