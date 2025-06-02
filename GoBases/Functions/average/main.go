package main

import "fmt"

func main() {

	var grade int

	var sum int

	fmt.Printf("Insira a quantidade de notas: ")

	fmt.Scanf("%d", &grade)

	for i := range grade {

		var value int

		fmt.Print("Insira a nota do aluno: ", i+1)

		fmt.Scan(&value)

		sum += value
	}

	media := Average(grade, sum)
	fmt.Printf("A média das notas é: %d\n", media)
}

func Average(grade int, value int) int {

	sum := value / grade

	return sum

}
