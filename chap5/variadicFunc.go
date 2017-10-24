package main

import "fmt"

// variadic funtion is one that can be called with vaiable no of arguments
// the type of FINAL agrument is proceded by elipsis, "..."
func sum(vals ...int) int {
	var total int
	for _, val := range vals {
		total += val
	}

	return total
}

func minus(mval ...int, pval ...int) {
	p, m := 0, 0
	for _, mv := range mval {
		m += mv
	}
	for _, pv := range pval {
		p += pv
	}
	fmt.Println(p, m)
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15
	fmt.Println(sum(1))             //1
	//minus(1, 1, 1, 1, 232, 3,2 ,3 2,3) // can only use ... with final parameter in list
}
