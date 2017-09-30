// Dup2 prints the count and text of lines that appears more thn once
// in the input. it reads from stdin or from  list of names files
package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if (len(files) == 0) {
		countlines(os.Stdin, counts)
	}
}

func countlines(f *os.File, counts map[string]int) {

}