package main

import "fmt"

func main() {

	var category string
	var hours int

	fmt.Println("Entre com as suas horas: ")
	fmt.Scanf("%d", &hours)

	fmt.Println("Entre com a sua categoria (A, B ou C): ")
	fmt.Scanf("%s", &category)

	res := salary(hours, category)

	fmt.Printf("Seu salário é de: %d\n", res)
}

func salary(hours int, category string) int {
	var baseSalary int

	switch category {
	case "A":
		baseSalary = 3000 * hours
		return baseSalary + baseSalary/2
	case "B":
		baseSalary = 1500 * hours
		return baseSalary + baseSalary/5
	case "C":
		return 1000 * hours
	default:
		return 0
	}
}
