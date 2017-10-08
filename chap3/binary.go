package main

import (
	"fmt"
)

func main() {
	var x uint8 = 1 << 1 | 1 << 5
	var y uint8 = 1 << 1 | 1 << 2
	// %b to print numbers binary digit
	fmt.Printf("x = %b\n", x)
	fmt.Printf("y = %b\n", y)
	fmt.Printf("z = %b\n", 4)

	// %08 to pad a result with zeros to exaclty 8 digit
	fmt.Printf("x = %08b\n", x)
	fmt.Printf("y = %08b\n", y)
	fmt.Printf("z = %08b\n", 4)


	fmt.Printf("%08b\n", x & y) // the intersection
	fmt.Printf("%08b\n", x | y) // the union
	fmt.Printf("%08b\n", x ^ y) // the symetric difference
	fmt.Printf("%08b\n", x &^ y) // the difference

	for i := uint(0); i < 8; i++ {
		if (x & (1 << i)) != 0 {
			fmt.Printf("%d ", i)
		}
	}
}
