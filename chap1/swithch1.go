package main

import (
	"fmt"
)
func coinflip()  string {
	return "tail"
}

func main() {
	switch coinflip() {
	case "tail":
		fmt.Println("got tail")
	case "head":
		fmt.Println("got head")
	default:
		fmt.Println("landed on edge!")
	}
}
