package main

import (
	"fmt"
)

func printhello(s string) string {
	return "hello " + s
}

var sayHi =  func(s string) string {return "Hi " + s}

func main() {
	fmt.Println(printhello("ram"))
	fmt.Println(sayHi("ram"))
}
