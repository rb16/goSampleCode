package main

import (
	"fmt"
)

//triple :- a deferred anonymous function can even change the value of enclosing function
// return to ist caller
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

// defer function run aftre return satements
func double(x int) (result int) {
	defer func() { fmt.Printf("Double of (%d) = %d\n", x, result) }()
	return x + 2
}

func main() {
	fmt.Println(double(2))
	fmt.Println(triple(2))
}
