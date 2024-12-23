package exx3

import (
	"fmt"
)

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

func (fte FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Salary: %d", fte.ID, fte.Name, fte.Salary)
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (pte PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Hourly Rate: %d, Hours Worked: %.1f", pte.ID, pte.Name, pte.HourlyRate, pte.HoursWorked)
}

type Company struct {
	Employees map[string]Employee
}

func NewCompany() *Company {
	return &Company{Employees: make(map[string]Employee)}
}

func (c *Company) AddEmployee(emp Employee) {
	c.Employees[fmt.Sprintf("%T-%v", emp, emp)] = emp
	fmt.Println("Employee added successfully.")
}

func (c *Company) ListEmployees() {
	for _, emp := range c.Employees {
		fmt.Println(emp.GetDetails())
	}
}
