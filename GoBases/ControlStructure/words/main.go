package main

import "fmt"

func main() {

	var a string

	fmt.Printf("Insira uma palavra: ")
	fmt.Scan(&a)

	count := 0

	for _, c := range a {
		fmt.Println(string(c))
		count++
	}

	fmt.Printf("Quantidade de letras: %d\n", count)

}
