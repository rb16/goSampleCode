package main

import (
	"fmt"
)

func signum(x int) int{
	// switch cases are evaluated from top to bottom
	// the [optional] default case match if none of the other cases matches
	// it can be any where
	switch {
	case x > 0 :
		return 1
	default:
		return 0
	case x < 0:
		return -1
	}
}

func main() {
	fmt.Println(signum(0))
	fmt.Println(signum(-1))
	fmt.Println(signum(1))
}
