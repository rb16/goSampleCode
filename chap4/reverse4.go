package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := [...]int{1, 2,3, 4, 5,6, 67,8}
	reverse(s[:])
	fmt.Println(s)
	reverse(s) //this is not slice(variable length, pointer to first element)
	fmt.Println(s) 
}
