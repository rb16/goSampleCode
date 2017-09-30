// Dup1 prints the text of each line that appears more than
// once in the standard inputm, presented by count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Note - Map has no order in golang
	for i, n := range counts {
		if n > 1 {
			//fmt.Println(n)
			fmt.Printf("%d\t%s\n", n, i)
		}
	}

	//After entering all the chacter or string
	// press "Ctrl + D" for the program to process
	// the input and end.
}
