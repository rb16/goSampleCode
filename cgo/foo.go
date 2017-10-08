package main

import "fmt"

/*
#include <stdio.h>
extern void ACFunction();
*/
import "C"

//export
func AGoFunction() {
	fmt.Println("A go function()")
}

func main() {
	C.ACFunction()
}
