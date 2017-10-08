package main

// this is the special comment section
// this piece of code is refecenced for go file

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello from stdio")
	// When you create a C string with C.CString (or any C memory allocation)
	// you must remember to free the memory when you're done with it by calling C.free.
	defer C.free(unsafe.Pointer(cs))
	C.myprint(cs)
}
