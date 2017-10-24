package main

import "fmt"

type point struct {
	X, Y int
}

type circle struct {
	Center point
	Radius int
}

type wheel struct {
	Circle circle
	Spoke  int
}

func main() {
	var w wheel
	w.Circle.Center.X = 12
	w.Circle.Center.Y = 13
	w.Spoke = 23
	w.Circle.Radius = 23
	fmt.Println(w)
}
