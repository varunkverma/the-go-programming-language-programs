package main

import (
	"fmt"
	"time"
)

// Employee stores and employee's info
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func (emp Employee) displayEmployeeInfo() {
	fmt.Println("ID: ", emp.ID)
	fmt.Println("Name: ", emp.Name)
	fmt.Println("Manager ID: ", emp.ManagerID)
	fmt.Println("Position : ", emp.Position)
	fmt.Println("Salary : ", emp.Salary)
}

func promote(emp *Employee, newPosition string) {
	emp.Position = newPosition
}

func increment(emp *Employee, newIncrement int) {
	(*emp).Salary += newIncrement
}

func main() {
	var dilbert Employee = Employee{
		ID:        10001,
		Name:      "Dilbert Shrut",
		Address:   "12th Hocking street",
		DoB:       time.Date(1994, time.April, 24, 11, 36, 59, 0, time.Local),
		Position:  "Software Engineer",
		Salary:    1000000,
		ManagerID: 10002,
	}
	dilbert.displayEmployeeInfo()
	fmt.Println()
	promote(&dilbert, "Senior Software Engineer")
	increment(&dilbert, 150000)
	dilbert.displayEmployeeInfo()
}
