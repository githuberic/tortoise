package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	Bob       time.Time
	Position  string
	Salary    int
	ManagerId int
}

func main() {
	var dilbert Employee
	dilbert.Name = "老刘"
	dilbert.Salary = 12999
	dilbert.Position = "Director"
	fmt.Println(dilbert)

	var employeeOfTheMonth *Employee = &dilbert
	(*employeeOfTheMonth).Position += "proactive team player"
	fmt.Println(*employeeOfTheMonth)
	employeeOfTheMonth.Position = "Senior director"
	fmt.Println(*employeeOfTheMonth)
}
