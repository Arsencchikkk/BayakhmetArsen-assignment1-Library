package main

import (
	"github.com/Arsencchikkk/BayakhmetArsen-assignment1-Library/exx3"
)

func main() {
	// Создаем экземпляр компании
	company := exx3.NewCompany()

	// Добавляем сотрудников
	company.AddEmployee(exx3.FullTimeEmployee{
		ID:     1,
		Name:   "John Doe",
		Salary: 50000,
	})

	company.AddEmployee(exx3.PartTimeEmployee{
		ID:          2,
		Name:        "Jane Smith",
		HourlyRate:  2000,
		HoursWorked: 20.5,
	})

	// Отображаем всех сотрудников
	company.ListEmployees()
}
