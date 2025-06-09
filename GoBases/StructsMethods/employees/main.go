package main

import "fmt"

type Person struct {
	Id          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	Id       int
	Position string
	Person   Person
}

func PrintEmployee(employees []Employee, name string) *Employee {
	for i := range employees {
		if employees[i].Person.Name == name {
			return &employees[i]
		}
	}
	return nil
}

func main() {
	employees := []Employee{
		{Id: 1, Position: "Manager", Person: Person{Id: 1, Name: "Alice", DateOfBirth: "1990-01-01"}},
		{Id: 2, Position: "Developer", Person: Person{Id: 2, Name: "Bob", DateOfBirth: "1992-02-02"}},
	}

	employee := PrintEmployee(employees, "Alice")
	if employee != nil {
		fmt.Println(employee)
	} else {
		fmt.Println("Employee not found")
	}
}
