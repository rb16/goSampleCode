package main

import "fmt"
import "math"

type Point struct{ X, Y float64 }

// A Path is a journey connecting the pints with sraight lines
type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// rule
// 1. All method of same types should have distinct name, but different types can have method name
//
func main() {
	// there is no conflict between the two declation of function called Distance, the first decrales a package-level function
	// call geometry.Distance, the second declares a method of the type Point, so its name is Point.Distance

	p := Point{X: 1, Y: 2} // "5", function call
	q := Point{4, 6}       // "5" method call

	fmt.Println(Distance(p, q))

	fmt.Println(p.Distance(q))

	prem := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(prem.Distance())
}
