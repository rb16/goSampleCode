package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	MAnagerID int
}

func main() {
	e := Employee{}
	fmt.Println(e)
}
