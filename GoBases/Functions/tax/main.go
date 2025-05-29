package main

import "fmt"

func main() {

	var salary float64

	fmt.Println("Entre com o seu salário: ")
	fmt.Scanf("%f", &salary)
	fmt.Println("O valor do seu imposto é:", tax(salary))

}

func tax(salary float64) float64 {
	if salary > 150000 {
		return salary * 0.27
	} else if salary > 50000 {
		return salary * 0.17
	} else {
		return 0
	}
}
