package main

import "fmt"

func main() {

	var age int

	var job int

	var salary int

	fmt.Printf("Insira sua idade: ")

	fmt.Scan(&age)

	fmt.Printf("Insira quanto tempo você está no seu atual emprego: ")

	fmt.Scan(&job)

	fmt.Printf("Insira seu salário: ")

	fmt.Scan(&salary)

	if age <= 22 {
		fmt.Printf("Você não está elegível ao empréstimo")
	} else if job <= 1 {
		fmt.Printf("Você não está elegível ao empréstimo")

	} else if salary <= 100000 {
		fmt.Printf("Você não está elegível ao empréstimo")
	} else {
		fmt.Printf("Empréstimo concendido")
	}
}
