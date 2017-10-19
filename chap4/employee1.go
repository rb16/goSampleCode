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
	ManagerID int
}

func EmployeeIDD(id int) Employee {

	return Employee{
		ID:        id,
		Name:      "my name",
		Address:   "my address",
		DoB:       time.Now(),
		Position:  "manager",
		Salary:    11111,
		ManagerID: 11,
	}
}
func EmployeeID(id int) *Employee {
	return &Employee{
		ID:        id,
		Name:      "my name",
		Address:   "my address",
		DoB:       time.Now(),
		Position:  "manager",
		Salary:    11111,
		ManagerID: 11,
	}
}
func main() {

	//e := EmployeeID(11)

	id := EmployeeIDD(11).ID
	// the assignment statement would not compile since the left hand side
	// would not identify a variable :(
	//TODO - POC
	EmployeeIDD(id).Name = "my new name"

	//fmt.Println(e)

	//p := EmployeeIDD(12)
	//fmt.Println(p)
	fmt.Println(EmployeeIDD(12).ID)

}
