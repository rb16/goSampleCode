package main

import "fmt"

type Point struct {
	X, Y int
}

//struct lierals
func main() {
	// this is the first form of struct literals
	// this is first one, requires that a value be specified by evert field by right order
	// It burdens the writer and reader to remember exactly what the fields are
	p := Point{1, 3}
	fmt.Println(p.Y, p.X)
}
