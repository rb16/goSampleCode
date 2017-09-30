package main

import (
	"fmt"
	"os"
)

func main() {
	var sep string
	var s string
	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}
	fmt.Println(s)
}
