// echo2 prints its commandline argument
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// string.Join is simple and efficient method
	// int the previous code we were using += to concanat string but += statement makes
	// a new string by concaneting by old string.
	// the old content is no longer in use, so they will be garbage collected
	fmt.Println(strings.Join(os.Args[1:], " "))
}
