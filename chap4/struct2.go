package main

import "fmt"

//embedded struct

type circle struct {
	X, Y, Radius int
}

type wheel struct {
	X, Y, Radius, Spokes int
}

func main() {
	var w wheel
	w.X = 8
	w.Y = 9
	w.Radius = 12
	w.Spokes = 15

	fmt.Println("wheel", w)
}
