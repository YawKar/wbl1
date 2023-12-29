package _24

import "math"

type Point struct {
	x, y float64
}

func MkPoint(x, y float64) Point {
	return Point{x, y}
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) EuclideanDistance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func (p *Point) ManhattanDistance(other *Point) float64 {
	return math.Abs(p.x-other.x) + math.Abs(p.y-other.y)
}
