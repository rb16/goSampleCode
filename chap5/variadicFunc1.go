package main

import (
	"fmt"
)

// although the ...int behaves same like slice but type of both function is different

func f(...int) {}
func g([]int)  {}

func main() {
	fmt.Printf("%T\n", f) // func f(...int) {}
	fmt.Printf("%T\n", g) // func g([]int)  {}
}
