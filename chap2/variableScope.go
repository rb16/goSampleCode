package main

import (
	"fmt"
	"os"
)

var cwd string
func init() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("os.Getwd got error: %v", err)
	}
	fmt.Printf("current working directory is = %s", cwd)
}

func main() {
	fmt.Println("hello")
}
