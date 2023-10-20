package iinv

import "math"

/*
This talks about interface seggeration
Like Volume is not meant for sqaure so that shouldn't implement
that interface
*/
type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

// No Need to impl volume on this
func (s Square) Volume() float64 {
	return 0
}

type Cube struct {
	Length float64
}

func (c Cube) Area() float64 {
	return math.Pow(c.Length, 2)
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Length, 3)
}

type figureI interface {
	Area() float64
	Volume() float64
}

func AreaTotal(figures ...figureI) float64 {
	var sum float64
	for _, s := range figures {
		sum += s.Area()
	}
	return sum
}

func AreaVolumeTotal(figures ...figureI) float64 {
	var sum float64
	for _, s := range figures {
		sum += s.Area() + s.Volume()
	}
	return sum
}
