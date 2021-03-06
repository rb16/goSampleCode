package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){
	counter := make(map[string]int)
	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counter[line]++
		}
	}
	for line, n := range counter {
		fmt.Printf("%d\t%s\n", n, line)
	}
}