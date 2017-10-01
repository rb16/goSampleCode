package main

import (
	"fmt"
)
// even though both have same underlying type, float64, they are not same type,
// so they can not be compared or combined in arithmetic expression

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

// for everry type T, there is a corresponding conversion operation T(x), that converts the value x to type T.
// In ny case a conversion necer fails at run time

// CToF for convert cesiud to fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9/5 +32)
}

// FtoC - function to convert Fahrenheit to Celsius
func  FToC(f Fahrenheit) Celsius{
	return Celsius((f - 32)* 5/9)
}

// type convertiosn - any compare compare a value of a named type to another of same type, or to a value of underlying type
// var f Fahrenheit
// var c Celsius
// fmt.Println(c == 0) //"true"
// fmt.Println(f >= 0) // true
// fmt.Println(c == f) // compile error: typr mismatch
// fmt.Println(c == Celsius(f)) // "ture"

func main() {
	fmt.Printf("%g\n", BoilingC - FreezingC) //100
	boilingF := CToF(BoilingC) //converting BoilingC which is of  type celsius
	fmt.Printf("%g\n", boilingF - CToF(FreezingC)) // first converting FreeZingC to Fahrenheit, finally both are same type
	fmt.Printf("%g\n", boilingF - FreezingC) // error: invalid operation: boilingF - FreezingC (mismatched types Fahrenheit and Celsius)
}
